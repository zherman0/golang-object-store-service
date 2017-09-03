package main


type Bucket struct {
	bucketName      string    `json:"name"`
	ListOfObj 	[]string  `json:"listofobj"`
}

type ObjectFile struct {
	FileName     string 	`json:"name"`
	FileData     string	`json:"data"`
}

type Buckets []Bucket
type ObjectFiles []ObjectFile