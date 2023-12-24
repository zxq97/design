package method

import "gorm.io/gen"

type LikeMethod interface {
	//sql(insert into @@table(obj_id, obj_type, uid, author, state) values (@objID, @objType, @uid, @author, 1) on deplicate key update state = 1)
	InsertLike(objID, uid, author int64, objType int32) error
	//sql(select uid, ctime from @@table where obj_id=@objID and obj_type=@objType and state=1 and id < (select id from @@table where obj_id=objID and obj_type=@objType and uid=@uid limit 1) order by id desc limit @limit)
	FindLikedUsers(objID, uid int64, objType int32, limit int8) ([]*gen.T, error)
	//sql(select obj_id, obj_type, uid, ctime from @@table where author=@author and obj_type=@objType and state=1 and id < (select id from @@table where (obj_id=@objID and obj_type=@objType and uid=@uid limit 1)) order by id desc limit @limit)
	FindRcvLikedRecord(objID, uid, author int64, objType int32, limit int8) ([]*gen.T, error)
	//sql(select obj_id, obj_type, author, ctime from @@table where uid=@uid and obj_type=@objType and state=1 and id < (select id from @@table where (obj_id=@objID and obj_type=@objType and uid=@author) limit 1) order by id desc limit @limit)
	FindPubLikeRecord(objID, uid, author int64, objType int32, limit int8) ([]*gen.T, error)
}

type LikeCountMethod interface {
	//sql(insert into @@table(obj_id, obj_type, count) values (@objID, @objType, @count) on duplicate key update count=count+@count)
	IncrByLikeCount(objID int64, objType, count int32) error
}
