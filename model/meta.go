package model

/*

"meta": {
  "data_version": "1.1.0",
  "created": "2020-07-06",
  "revision": 1
}

*/

type MetaInfo struct {
	DataVersion string `json:"data_version"` // Version of the data format (required)
	Created     string `json:"created"`      // Date when the data file was created (required, format: YYYY-MM-DD)
	Revision    int    `json:"revision"`     // Revision number of the data file (required)
}
