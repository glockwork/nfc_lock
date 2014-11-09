package helpers

import (
    "encoding/hex"
//    "encoding/binary"
    "github.com/fuzxxl/freefare/0.3/freefare"
)


func String2aeskey(keydata_str string) (*freefare.DESFireKey, error) {
    keydata := new([16]byte)
    to_keydata, err := hex.DecodeString(keydata_str)
    if err != nil {
        key := freefare.NewDESFireAESKey(*keydata, 0)
        return key, err
    }
    copy(keydata[0:], to_keydata)
    key := freefare.NewDESFireAESKey(*keydata, 0)
    return key,nil
}

func Bytes2aeskey(source []byte) (*freefare.DESFireKey) {
    keydata := new([16]byte)
    copy(keydata[0:], source)
    key := freefare.NewDESFireAESKey(*keydata, 0)
    return key
}

func String2byte(source string) (byte, error) {
    bytearray, err := hex.DecodeString(source)
    if err != nil {
        return 0x0, err
    }
    return bytearray[0], nil
}

func Applicationsettings(accesskey byte, frozen, req_auth_fileops, req_auth_dir, allow_master_key_chg bool) byte {
    ret := byte(0)
    ret |= accesskey << 4
    if (frozen) {
        ret |= 1 << 3
    }
    if (req_auth_fileops) {
        ret |= 1 << 2
    }
    if (req_auth_dir) {
        ret |= 1 << 1;
    }
    if (allow_master_key_chg) {
        ret |= 1;
    }
    return ret
}
