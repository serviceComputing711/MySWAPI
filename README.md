# <center>MySWAPI设计文档<center>
RESTful API SWAPI written in GO
----------------------------------

## 1. 数据库—数据收集

### 1.1 设计思路

> 由于选择了`StarWar`作为样例。因此使用`轮询请求数据`的方法获取数据。
> 数据由6个部分组成：
> - 1.films
> - 2.person
> - 3.species
> - 4.planets
> - 5.vehicles
> - 6.starships
> 这些数据层层嵌套，使用大量的`URL`关联其他数据。
> 由文档所提供的`Go`语言的源代码实现了访问获取数据。因此基本思路为反复调用，逐条获取数据。

------------------------------
### 1.2 代码设计

> 源码中，通过调用一个`Client`的`Jump`来取得数据，修改函数。
> 修改`Jump`函数，将得到的结构体直接返回。
```
func dump(data interface{}, err error) (mes interface{}) {
	if err != nil {
		return
	}
	return data
}
```

> 这样，就可以得到结构，再通过转化的方式，得到`[]byte`形式的`Json`。

```
		film := dump(c.Film(i))
		//transport json to data
		data, _ := json.Marshal(film)
		b := []byte(data)
```

---------------------------------

## 2. 数据库—数据处理

----------------------

### 2.1 数据库设计思路

> 数据库实现了几个基本功能，分别为：查询，增加条目，删除条目。
> 由于使用`boltdb`，因此存储的时候只能够存储`key-value`对，因此在设计时，为每个数据类型设计一个`bucket`，将`ID`作为`key`，`Json`格式的数据作为`value`，进行存储。
> 为了方便操作，封装起来。

```
type SWDB struct {
	db *bolt.DB
}
```

--------------------

### 2.2 数据库初始化

> 初始化桶，由于涉及数据的修改，因此使用操作事务。
```
	//use operation bussiness
	err := m.db.Update(func(tx *bolt.Tx) error {
		//create buckets
		_, err := tx.CreateBucketIfNotExists([]byte("films"))
		_, err = tx.CreateBucketIfNotExists([]byte("people"))
		_, err = tx.CreateBucketIfNotExists([]byte("planets"))
		_, err = tx.CreateBucketIfNotExists([]byte("species"))
		_, err = tx.CreateBucketIfNotExists([]byte("starships"))
		_, err = tx.CreateBucketIfNotExists([]byte("vehicles"))
		_, err = tx.CreateBucketIfNotExists([]byte("user"))
		return err
	})
	return err
```

> 初始化数据。
> 由于数据的初始化需要借助一条一条的将数据加入到对应的`bucket`里，因此，在数据获取的时候完成将数据放入数据库的操作：

```
		//add information to db
		db.AddObj("films", []byte(strconv.Itoa(i)), b)
```

--------------------

### 2.3 增加模块

> 涉及数据的修改，因此使用操作事务。
> 有两个部分，第一个部分为加入常规的数据，第二个部分为加入登陆信息，即`token-UserName`。
> 对于直接找到对应的`bucket`，插入即可。

```
/*
**	add information to bucket
**	has 3 arguement:
**	buc: buckets type, as films, people, etc
**	key: object id
**	value: object information as Json
**	return wrong message
*/
func (m *SWDB) AddObj(buc string, key []byte, value []byte) (error) {
	//use operation bussiness
	err := m.db.Update(func(tx *bolt.Tx) error {
		//find the buckets
		b := tx.Bucket([]byte(buc))
		//add information; has error, return wrong message
		return b.Put(key, value)
	})
	// fmt.Println("add a new information")
	return err
}
```

> 插入登陆信息大同小异，不赘述。

--------------------------

### 2.4 删除数据

> 涉及数据的修改，因此使用操作事务。
> 直接找到对应的`bucket`对应的条目，删除即可。

```
/*
**	delete a object from database
**	has 2 arguement
**	buc: buckets type, as films, people, etc
**	key: object id
**	if has error, return wrong message
*/
func (m *SWDB) DeleteObj(buc string, key []byte) (error) {
	//use operation bussiness
	err := m.db.Update(func(tx *bolt.Tx) error {
		//find the buckets
		b := tx.Bucket([]byte(buc))
		//delete element whose id == key
		return b.Delete(key)
	})
	return err
}
```

----------------------------

### 2.5 检索模块

> 检索模块分多个部分，主要有：查询是否存在；按ID检索并返回信息；按姓名检索并返回信息；按页检索并返回信息。
> 由于没有涉及数据的修改，因此使用查看事务。
> 查询是否存在：

```
/*
**	check object exsist or not by name
**	has 2 arguement
**	buc: buckets type, as films, people, etc
**	key: object name
**	return bool
*/
func (m *SWDB) HasObj (buc string, name []byte) (bool) {
	flag := false
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(buc))
        c := b.Cursor()
        var p Person
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				flag = true
				break
			}
		}
        return nil
    })
    if (flag) {
    	return true
    } else {
    	return false
    }
}
```

> 按ID检索

```
/*
**	search object by ID
**	has 2 arguement
**	buc: buckets type, as films, people, etc
**	key: object ID
**	return object by Json
*/
func (m *SWDB) SearchByID (buc string, key []byte) ([]byte) {
	var v []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(buc))
        v = b.Get(key)
        return nil
    })
    return v
}
```

> 按姓名检索，由于数据存储的时候存储了一整个`Json`，因此需要将其转换为结构体并添加获得姓名的函数。获得姓名函数如下：

```
func (p *Film) getName() (string){
	return p.Title
}
```

> 每一个类都有对应的函数，不赘述。
> 检索函数如下：

```
/*
**	search film by Name
**	has 1 arguement
**	key: object Name
**	return object by Json
*/
func (m *SWDB) SearchFilmByName (name []byte) ([]byte) {
	var res []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("films"))
        c := b.Cursor()
        var p Film
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				res = v
				break
			}
		}
        return nil
    })
    return res
}
```

> 按页检索：每页大小设置为5，从头开始遍历，访问到页区间则加入到输出的结果中。
> 由于涉及到`[]byte`的拼接，因此，采用`Buffer`流的方法：

```
/*
**	search information by page
**	has 2 arguement
**	buc: buckets, like people, films, etc
**	page: page number
*/
func (m *SWDB) SearchByPage (buc string, page int) ([]byte) {
	var buffer bytes.Buffer
	endNum := 5 * page
	startNum := endNum - 5
	currentNum := 0
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(buc))
        c := b.Cursor()
	    for k, v := c.First(); k != nil; k, v = c.Next() {
	    	if currentNum < endNum && currentNum >= startNum {
	    		buffer.Write(v)
	    	}
			currentNum++
			if currentNum == endNum {
				break
			}
		}
        return nil
    })
    return buffer.Bytes()
}
```
