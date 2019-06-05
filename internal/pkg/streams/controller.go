package streams

type StorageReader interface{
	ReadByID(idStr string) (*Stream, error)
}

func ReadStream(streamID string) (*Stream, error) {
	dbReader, err :=  NewFileReader()
	if err != nil {
		return nil, err
	}
	record, err := dbReader.ReadByID(streamID)
	if err != nil {
		return nil, err
	}
	if record == nil {
		// there is no record
		return nil, nil
	}	

	stream := Stream{
		ID: record.ID,
		StreamUrl: record.StreamUrl,
		Captions: record.Captions,
		Ads: nil,
	}
	if stream.Ads, err = fetchAds(streamID); err != nil {
		return nil, err
	}
	
	return &stream, nil
}
