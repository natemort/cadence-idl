// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/api/v1/tasklist.proto

package apiv1

var yarpcFileDescriptorClosure216fa006947e00a0 = [][]byte{
	// uber/cadence/api/v1/tasklist.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdb, 0x6e, 0xdb, 0x36,
		0x18, 0x9e, 0x7c, 0x68, 0x9d, 0xdf, 0x4d, 0xa2, 0xb2, 0x49, 0x63, 0xbb, 0xed, 0xe6, 0xfa, 0xa2,
		0xc8, 0x8a, 0x4d, 0x46, 0xb2, 0x0d, 0x18, 0xb6, 0xa1, 0xab, 0x13, 0x1b, 0xad, 0x10, 0x27, 0x35,
		0x64, 0xb5, 0x43, 0x07, 0x0c, 0x02, 0x2d, 0xb1, 0x0e, 0x67, 0x49, 0x14, 0x44, 0xca, 0xae, 0x6f,
		0xf6, 0x18, 0xbb, 0xdb, 0x8b, 0xec, 0x1d, 0xf6, 0x4e, 0x03, 0x29, 0xd9, 0xf5, 0x41, 0x0d, 0xd6,
		0x8b, 0xdd, 0x99, 0xff, 0xc7, 0x8f, 0xdf, 0x7f, 0xb6, 0xa0, 0x95, 0x8c, 0x48, 0xdc, 0x76, 0xb1,
		0x47, 0x42, 0x97, 0xb4, 0x71, 0x44, 0xdb, 0xd3, 0x93, 0xb6, 0xc0, 0x7c, 0xe2, 0x53, 0x2e, 0x8c,
		0x28, 0x66, 0x82, 0xa1, 0x7b, 0xf2, 0x8e, 0x91, 0xdd, 0x31, 0x70, 0x44, 0x8d, 0xe9, 0x49, 0xe3,
		0xf3, 0x31, 0x63, 0x63, 0x9f, 0xb4, 0xd5, 0x95, 0x51, 0xf2, 0xae, 0xed, 0x25, 0x31, 0x16, 0x94,
		0x85, 0x29, 0xa9, 0xf1, 0xc5, 0x26, 0x2e, 0x68, 0x40, 0xb8, 0xc0, 0x41, 0x94, 0x5d, 0xd8, 0x7a,
		0x60, 0x16, 0xe3, 0x28, 0x22, 0x31, 0x4f, 0xf1, 0xd6, 0x6b, 0xa8, 0xd8, 0x98, 0x4f, 0xfa, 0x94,
		0x0b, 0x84, 0xa0, 0x14, 0xe2, 0x80, 0xd4, 0xb4, 0xa6, 0x76, 0xbc, 0x63, 0xa9, 0xdf, 0xe8, 0x3b,
		0x28, 0x4d, 0x68, 0xe8, 0xd5, 0x0a, 0x4d, 0xed, 0x78, 0xef, 0xf4, 0xb1, 0x91, 0xe3, 0xa4, 0xb1,
		0x78, 0xe0, 0x82, 0x86, 0x9e, 0xa5, 0xae, 0xb7, 0x30, 0xe8, 0x0b, 0xeb, 0x25, 0x11, 0xd8, 0xc3,
		0x02, 0xa3, 0x4b, 0x38, 0x08, 0xf0, 0x7b, 0x47, 0x86, 0xcd, 0x9d, 0x88, 0xc4, 0x0e, 0x27, 0x2e,
		0x0b, 0x3d, 0x25, 0x57, 0x3d, 0x7d, 0x68, 0xa4, 0x9e, 0x1a, 0x0b, 0x4f, 0x8d, 0x2e, 0x4b, 0x46,
		0x3e, 0x79, 0x83, 0xfd, 0x84, 0x58, 0x77, 0x03, 0xfc, 0x5e, 0x3e, 0xc8, 0x07, 0x24, 0x1e, 0x2a,
		0x5a, 0xeb, 0x35, 0xd4, 0x17, 0x12, 0x03, 0x1c, 0x0b, 0x2a, 0xb3, 0xb2, 0xd4, 0xd2, 0xa1, 0x38,
		0x21, 0xf3, 0x2c, 0x12, 0xf9, 0x13, 0x3d, 0x81, 0x7d, 0x36, 0x0b, 0x49, 0xec, 0x5c, 0x33, 0x2e,
		0x1c, 0x15, 0x67, 0x41, 0xa1, 0xbb, 0xca, 0xfc, 0x92, 0x71, 0x71, 0x85, 0x03, 0xd2, 0x9a, 0xc0,
		0xa1, 0xc9, 0x99, 0xaf, 0x92, 0xfc, 0x22, 0x66, 0x49, 0x74, 0x49, 0x44, 0x4c, 0x5d, 0x8e, 0xda,
		0x70, 0x10, 0x92, 0x59, 0xbe, 0xfb, 0x9a, 0x75, 0x37, 0x24, 0xb3, 0x75, 0x07, 0xd1, 0x63, 0xb8,
		0x13, 0x31, 0xdf, 0x27, 0xb1, 0xe3, 0xb2, 0x24, 0x14, 0x4a, 0xae, 0x68, 0x55, 0x53, 0xdb, 0xb9,
		0x34, 0xb5, 0xfe, 0x2a, 0xc1, 0xde, 0x22, 0x88, 0xa1, 0xc0, 0x22, 0xe1, 0xe8, 0x2b, 0x40, 0x23,
		0xec, 0x4e, 0x7c, 0x36, 0x4e, 0x69, 0xce, 0x35, 0x0d, 0x85, 0x12, 0x29, 0x5a, 0x7a, 0x86, 0x28,
		0xf2, 0x4b, 0x1a, 0x0a, 0xf4, 0x08, 0x20, 0x26, 0xd8, 0x73, 0x7c, 0x32, 0x25, 0x7e, 0xa6, 0xb0,
		0x23, 0x2d, 0x7d, 0x69, 0x40, 0x0f, 0x60, 0x07, 0xbb, 0x93, 0x0c, 0x2d, 0x2a, 0xb4, 0x82, 0xdd,
		0x49, 0x0a, 0x3e, 0x81, 0xfd, 0x18, 0x0b, 0xb2, 0x1a, 0x4b, 0x49, 0xc5, 0xb2, 0x2b, 0xcd, 0x1f,
		0xe2, 0xe8, 0xc2, 0xae, 0x0c, 0xda, 0xa1, 0x9e, 0x33, 0xf2, 0x99, 0x3b, 0xa9, 0x95, 0x55, 0xc1,
		0x9a, 0x1f, 0xed, 0x05, 0xb3, 0x7b, 0x26, 0xef, 0x59, 0x55, 0x49, 0x33, 0x3d, 0x75, 0x40, 0x53,
		0x38, 0xa2, 0x8b, 0xbc, 0x3a, 0x63, 0x99, 0x58, 0x27, 0x48, 0x33, 0x5b, 0xbb, 0xd5, 0x2c, 0x1e,
		0x57, 0x4f, 0x9f, 0xdd, 0xd8, 0x5b, 0x69, 0x76, 0x8c, 0xdc, 0xd2, 0xf4, 0x42, 0x11, 0xcf, 0xad,
		0x43, 0xfa, 0x49, 0x65, 0xbb, 0xfd, 0xb1, 0xb2, 0x1d, 0x40, 0x99, 0x04, 0x91, 0x98, 0xd7, 0x2a,
		0x4d, 0xed, 0xb8, 0x62, 0xa5, 0x87, 0x86, 0x80, 0xc6, 0xc7, 0xb5, 0x73, 0xda, 0xed, 0x39, 0x94,
		0xa7, 0xb2, 0x73, 0x55, 0x4d, 0xaa, 0xa7, 0x4f, 0x73, 0x83, 0xcb, 0x7d, 0xd1, 0x4a, 0x89, 0x3f,
		0x14, 0xbe, 0xd7, 0x5a, 0x3f, 0x43, 0x75, 0x25, 0xa1, 0xa8, 0x0e, 0x15, 0x2e, 0x70, 0x2c, 0x1c,
		0xea, 0x65, 0x1d, 0x71, 0x5b, 0x9d, 0x4d, 0x0f, 0x1d, 0xc2, 0x2d, 0x12, 0x7a, 0x12, 0x48, 0x9b,
		0xa0, 0x4c, 0x42, 0xcf, 0xf4, 0x5a, 0x7f, 0x6a, 0x00, 0x03, 0xd5, 0x70, 0x66, 0xf8, 0x8e, 0xa1,
		0x2e, 0xe8, 0x3e, 0xe6, 0xc2, 0xc1, 0xae, 0x4b, 0x38, 0x77, 0xe4, 0xb2, 0xc8, 0xc6, 0xaf, 0xb1,
		0x35, 0x7e, 0xf6, 0x62, 0x93, 0x58, 0x7b, 0x92, 0xd3, 0x51, 0x14, 0x69, 0x44, 0x0d, 0xa8, 0x50,
		0x8f, 0x84, 0x82, 0x8a, 0x79, 0x36, 0x43, 0xcb, 0x73, 0x5e, 0x53, 0x15, 0x73, 0x9a, 0xaa, 0xf5,
		0xb7, 0x06, 0xf5, 0xa1, 0xa0, 0xee, 0x64, 0xde, 0x7b, 0x4f, 0xdc, 0x44, 0x26, 0xa1, 0x23, 0x44,
		0x4c, 0x47, 0x89, 0x20, 0x1c, 0xbd, 0x00, 0x7d, 0xc6, 0xe2, 0x09, 0x89, 0x55, 0xdd, 0x1c, 0xb9,
		0x25, 0x33, 0x3f, 0x1f, 0xdd, 0xd8, 0x25, 0xd6, 0x5e, 0x4a, 0x5b, 0xae, 0x34, 0x1b, 0xea, 0xdc,
		0xbd, 0x26, 0x5e, 0xe2, 0x13, 0x47, 0x30, 0x27, 0xcd, 0x9e, 0x0c, 0x9b, 0x25, 0x22, 0x2b, 0x4d,
		0x7d, 0x7b, 0xf1, 0x64, 0x3b, 0xd6, 0xba, 0xbf, 0xe0, 0xda, 0x6c, 0x28, 0x99, 0x76, 0x4a, 0x6c,
		0x3d, 0x83, 0xbb, 0x5b, 0xab, 0x07, 0x7d, 0x09, 0xfa, 0x46, 0x83, 0xf3, 0x9a, 0xd6, 0x2c, 0x1e,
		0xef, 0x58, 0xfb, 0xeb, 0x9d, 0xc9, 0x5b, 0xff, 0x94, 0xe0, 0x68, 0xeb, 0x81, 0x73, 0x16, 0xbe,
		0xa3, 0x63, 0x54, 0x83, 0xdb, 0x53, 0x12, 0x73, 0xca, 0xc2, 0x45, 0x89, 0xb3, 0x23, 0x3a, 0x85,
		0x7b, 0x61, 0x12, 0x38, 0x6a, 0xde, 0xa3, 0x05, 0x8b, 0xab, 0x28, 0xca, 0x67, 0x85, 0x9a, 0x6c,
		0xe6, 0x24, 0xb0, 0x08, 0xf6, 0x96, 0x4f, 0x72, 0xf4, 0x2d, 0x1c, 0x48, 0xce, 0x2c, 0xa6, 0xb2,
		0x26, 0x1f, 0x48, 0xc5, 0x25, 0x09, 0x85, 0x49, 0xf0, 0x8b, 0x84, 0x57, 0x58, 0x14, 0xf6, 0x37,
		0x55, 0x4a, 0x6a, 0x46, 0x9f, 0xdf, 0x98, 0xfd, 0x8d, 0x50, 0x8c, 0x75, 0x5f, 0xd2, 0x29, 0xdd,
		0x8b, 0xd7, 0x1d, 0xf4, 0x41, 0xdf, 0x72, 0xae, 0xac, 0xb4, 0x3a, 0x9f, 0xa4, 0xb5, 0x11, 0x42,
		0x2a, 0xb6, 0x3f, 0x5b, 0xb7, 0x36, 0x28, 0xdc, 0xcb, 0x71, 0x6a, 0x75, 0x7c, 0xcb, 0xe9, 0xf8,
		0xfe, 0xb4, 0x3e, 0xbe, 0x4f, 0xfe, 0x9b, 0x2f, 0x2b, 0xa3, 0xdb, 0xf8, 0x1d, 0x0e, 0xf2, 0x7c,
		0xfa, 0x3f, 0xb4, 0x9e, 0xfe, 0x01, 0x77, 0x56, 0xff, 0x83, 0x51, 0x03, 0xee, 0xdb, 0x9d, 0xe1,
		0x85, 0xd3, 0x37, 0x87, 0xb6, 0x73, 0x61, 0x5e, 0x75, 0x1d, 0xf3, 0xea, 0x4d, 0xa7, 0x6f, 0x76,
		0xf5, 0xcf, 0x50, 0x1d, 0x0e, 0x37, 0xb0, 0xab, 0x57, 0xd6, 0x65, 0xa7, 0xaf, 0x6b, 0x39, 0xd0,
		0xd0, 0x36, 0xcf, 0x2f, 0xde, 0xea, 0x05, 0xf4, 0x10, 0x6a, 0x1b, 0x50, 0x6f, 0xf0, 0xb2, 0x77,
		0xd9, 0xb3, 0x3a, 0x7d, 0xbd, 0xf8, 0xd4, 0xfb, 0xa0, 0x6f, 0xcf, 0x23, 0xb2, 0xae, 0x6f, 0xbf,
		0x1d, 0xf4, 0x56, 0xf4, 0x1f, 0xc0, 0xd1, 0x06, 0xd6, 0xed, 0x9d, 0x9b, 0x43, 0xf3, 0xd5, 0x95,
		0xae, 0xe5, 0x80, 0x9d, 0x73, 0xdb, 0x7c, 0x63, 0xda, 0x6f, 0xf5, 0xc2, 0xd9, 0x6f, 0x70, 0xe4,
		0xb2, 0x20, 0x2f, 0x3b, 0x67, 0xbb, 0xcb, 0xf4, 0xc8, 0x19, 0x1e, 0x68, 0xbf, 0x9e, 0x8c, 0xa9,
		0xb8, 0x4e, 0x46, 0x86, 0xcb, 0x82, 0xf6, 0xea, 0xb7, 0xd7, 0xd7, 0xd4, 0xf3, 0xdb, 0x63, 0x96,
		0x7e, 0x0e, 0x65, 0x1f, 0x62, 0x3f, 0xe2, 0x88, 0x4e, 0x4f, 0x46, 0xb7, 0x94, 0xed, 0x9b, 0x7f,
		0x03, 0x00, 0x00, 0xff, 0xff, 0xea, 0xcc, 0x39, 0x04, 0xac, 0x09, 0x00, 0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
	// google/protobuf/timestamp.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d,
		0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x03, 0x0b, 0x09, 0xf1, 0x43, 0x14, 0xe8, 0xc1, 0x14, 0x28,
		0x59, 0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5,
		0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89,
		0x79, 0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x53, 0x2b, 0x23, 0x97,
		0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xa1, 0x4e, 0x7c, 0x70, 0x23, 0x03, 0x40, 0x42, 0x01, 0x8c,
		0x51, 0x46, 0x50, 0x25, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0x48, 0x6e,
		0xac, 0x2c, 0x48, 0x2d, 0xd6, 0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0x43, 0xb8, 0xb7, 0x20, 0xe9, 0x07,
		0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10, 0xdd, 0x01, 0x50,
		0x2d, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x20, 0x0d, 0x21, 0x20, 0xbd, 0x49, 0x6c, 0x60, 0xb3,
		0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x65, 0xce, 0x7d, 0xff, 0x00, 0x00, 0x00,
	},
	// google/protobuf/wrappers.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c,
		0x28, 0x48, 0x2d, 0x2a, 0xd6, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0xca,
		0x5c, 0xdc, 0x2e, 0xf9, 0xa5, 0x49, 0x39, 0xa9, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x42, 0x22, 0x5c,
		0xac, 0x65, 0x20, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x84, 0xa3, 0xa4, 0xc4, 0xc5,
		0xe5, 0x96, 0x93, 0x9f, 0x58, 0x82, 0x45, 0x0d, 0x13, 0x92, 0x1a, 0xcf, 0xbc, 0x12, 0x33, 0x13,
		0x2c, 0x6a, 0x98, 0x61, 0x6a, 0x94, 0xb9, 0xb8, 0x43, 0x71, 0x29, 0x62, 0x41, 0x35, 0xc8, 0xd8,
		0x08, 0x8b, 0x1a, 0x56, 0x34, 0x83, 0xb0, 0x2a, 0xe2, 0x85, 0x29, 0x52, 0xe4, 0xe2, 0x74, 0xca,
		0xcf, 0xcf, 0xc1, 0xa2, 0x84, 0x03, 0xc9, 0x9c, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x2c, 0x8a,
		0x38, 0x91, 0x1c, 0xe4, 0x54, 0x59, 0x92, 0x5a, 0x8c, 0x45, 0x0d, 0x0f, 0x54, 0x8d, 0x53, 0x33,
		0x23, 0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x5a, 0xf0, 0x3a, 0xf1, 0x86, 0x43, 0xc3, 0x3f, 0x00,
		0x24, 0x12, 0xc0, 0x18, 0x65, 0x08, 0x55, 0x91, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f,
		0x94, 0x8e, 0x88, 0xab, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xfd, 0xec, 0xbc, 0xfc, 0xf2, 0x3c, 0x78,
		0xbc, 0x15, 0x24, 0xfd, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce,
		0x1d, 0xa2, 0x39, 0x00, 0xaa, 0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4,
		0x35, 0x89, 0x0d, 0x6c, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x92, 0x48, 0x30, 0x06,
		0x02, 0x00, 0x00,
	},
}
