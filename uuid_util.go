package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	uuid "github.com/gofrs/uuid"
)

// UUIDString UUIDString
func UUIDString(id uuid.UUID) string {
	return hex.EncodeToString(id.Bytes())
}

// UUIDGen UUIDGen
// b9be9a09-7117-4bdb-9c6d-737269c86480
func UUIDGen() string {
	return UUIDV4Gen()
}

// UUIDV1Gen UUIDV1Gen
// b9be9a09-7117-4bdb-9c6d-737269c86480
func UUIDV1Gen() string {
	return uuid.Must(uuid.NewV1()).String()
}

// UUIDV4Gen UUIDGen
// b9be9a09-7117-4bdb-9c6d-737269c86480
func UUIDV4Gen() string {
	return uuid.Must(uuid.NewV4()).String()
}

// UUIDStringGen UUIDStringGen
// 5106b5e58ee44f74a5d49e779dcf7f57
func UUIDStringGen() string {
	uuid := uuid.Must(uuid.NewV4())
	return UUIDString(uuid)
}

// UUIDFromString UUIDFromString
func UUIDFromString(value string) (uuid.UUID, error) {
	uid, _ := uuid.FromString(value)

	// if nil != err {
	// 	return , err
	// }

	return uid, nil
}

// SettlementID 根据规则生成 UUID
func SettlementID(id, modifier string) string {
	h := md5.New()
	io.WriteString(h, id)
	io.WriteString(h, modifier)
	sum := h.Sum(nil)
	sum[6] = (sum[6] & 0x0f) | 0x30
	sum[8] = (sum[8] & 0x3f) | 0x80
	return uuid.FromBytesOrNil(sum).String()
}
