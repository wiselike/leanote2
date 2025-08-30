package service

import (
	"github.com/wiselike/leanote2/app/info"
	//	. "github.com/wiselike/leanote2/app/lea"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote2/app/db"
)

const IMAGE_TYPE = 0

type AlbumService struct {
}

// add album
func (this *AlbumService) AddAlbum(album info.Album) bool {
	album.CreatedTime = time.Now()
	album.Type = IMAGE_TYPE
	return db.Insert(db.Albums, album)
}

// get albums
func (this *AlbumService) GetAlbums(userId string) []info.Album {
	albums := []info.Album{}
	db.ListByQ(db.Albums, bson.M{"UserId": bson.ObjectIdHex(userId)}, &albums)
	return albums
}

// get album by id
func (this *AlbumService) GetAlbumById(userId, albumId string) *info.Album {
	album := &info.Album{}
	if albumId == "" || albumId == "null" {
		return album
	}
	db.GetByIdAndUserId(db.Albums, albumId, userId, album)
	return album
}

// delete album
// presupposition: has no images under this ablum
func (this *AlbumService) DeleteAlbum(userId, albumId string) (bool, string) {
	if db.Count(db.Files, bson.M{"AlbumId": bson.ObjectIdHex(albumId),
		"UserId": bson.ObjectIdHex(userId),
	}) == 0 {
		return db.DeleteByIdAndUserId(db.Albums, albumId, userId), ""
	}
	return false, "has images"
}

// update album name
func (this *AlbumService) UpdateAlbum(albumId, userId, name string) bool {
	return db.UpdateByIdAndUserIdField(db.Albums, albumId, userId, "Name", name)
}
