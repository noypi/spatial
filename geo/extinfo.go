package geo

func (this *SpatialGeo) SetExtInfo(category uint8, id []byte, v interface{}) error {
	return this.db.SetExtInfo(category, id, v)
}

func (this *SpatialGeo) GetExtInfo(category uint8, id []byte) (v interface{}, err error) {
	return this.db.GetExtInfo(category, id)
}
