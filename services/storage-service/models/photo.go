package models

import (
	"time"
)

type Photo struct {
	//TODO: binary UUID encoding
	ID          string    `bson:"_id"`
	UserID      uint64    `bson:"user_id"`
	Kind        PhotoKind `bson:"kind"`
	FileName    string    `bson:"file_name"`
	FileSize    int64     `bson:"file_size"`
	ContentType string    `bson:"content_type"`
	UploadedAt  time.Time `bson:"uploaded_at"`
	Metadata    Metadata  `bson:"metadata,inline,omitempty"`
}

/* someday later
type MongoUUID struct {
	uuid.UUID
}

func NewMongoUUID(uuid uuid.UUID) MongoUUID {
	return MongoUUID{UUID: uuid}
}

func (u *MongoUUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	binaryUUID := primitive.Binary{
		Subtype: 0x04,
		Data:    u.UUID[:],
	}
	return bson.MarshalValue(binaryUUID)
}

func (u *MongoUUID) UnmarshalBSONValue(bt bsontype.Type, data []byte) error {
	if bt != bson.TypeBinary {
		return fmt.Errorf("expected bson type '%v', got '%v'", bson.TypeBinary, bt)
	}

	var binaryUUID primitive.Binary
	if err := bson.UnmarshalValue(bt, data, &binaryUUID); err != nil {
		return err
	}

	if binaryUUID.Subtype != 0x04 {
		return fmt.Errorf("expected bson type '%v', got '%v'", bson.TypeBinary, bt)
	}

	parsedUUID, err := uuid.FromBytes(binaryUUID.Data)
	if err != nil {
		return err
	}

	u.UUID = parsedUUID
	return nil
}
*/

type PhotoKind string

const (
	PhotoKindMedia      PhotoKind = "media"
	PhotoKindScreenshot PhotoKind = "screenshot"
	PhotoKindMeme       PhotoKind = "meme"
)
