package main

import (
	"C"
)

import (
	"fmt"

	. "github.com/windosx/face-engine/v2"
	"github.com/windosx/face-engine/v2/util"
)

//ArcSoft Face SDK 结合数据库的使用 https://github.com/tz-byte/arcface-with-database

var width, height = util.GetImageWidthAndHeight("./hj1.jpg")
var imageData = util.GetResizedBGR("./hj1.jpg")

func main() {
	AppId := "9z4iCPjUc3eKGDo5NDmLnZj6yzmyEypN7tKfv54ZLzSv"  // 官网申请的APP_ID
	SdkKey := "3E6L5TGV41JK5xGwLhAzSkdBvRiL7k59j41Ut4iy6xnX" // 官网申请的SDK_KEY

	// 激活SDK
	if err := Activation(AppId, SdkKey); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 初始化引擎
	engine, err := NewFaceEngine(DetectModeImage,
		OrientPriority0,
		12,
		50,
		EnableFaceDetect|EnableFaceRecognition|EnableFace3DAngle|EnableLiveness|EnableIRLiveness|EnableAge|EnableGender)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 检测人脸
	info, err := engine.ASFDetectFaces(width-width%4, height, ColorFormatBGR24, imageData)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("info: %v\n", info)
	// fmt.Printf("info.FaceID: %v\n", &info.FaceID)

	faceInfo, err := engine.DetectFaces(width-width%4, height, ColorFormatBGR24, imageData)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("faceInfo: %v\n", faceInfo)
	// fmt.Printf("faceInfo.FaceID: %v\n", &faceInfo.FaceID)

	singleFaceInfo := &SingleFaceInfo{
		FaceRect:   faceInfo.FaceRect[0],
		FaceOrient: faceInfo.FaceOrient[0],
	}

	ASFFaceFeature, faceFeature, err := engine.ASFFaceFeatureExtract(width-width%4, height, ColorFormatBGR24, imageData, singleFaceInfo)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("ASFFaceFeature: %v\n", ASFFaceFeature)
	fmt.Printf("faceFeature: %v\n", faceFeature)
	fmt.Printf("faceFeature.Feature->", faceFeature.Feature)
	// fmt.Printf("faceToken.feature: %v\n", (*[]int32)(unsafe.Pointer(faceToken.feature))[:])
	// fmt.Printf("faceToken.feature: %v\n", unsafe.Pointer(faceToken.feature))

	// 处理人脸数据
	if err = engine.Process(width-width%4, height, ColorFormatBGR24, imageData, info, EnableAge|EnableGender|EnableFace3DAngle|EnableLiveness); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 获取年龄
	ageInfo, err := engine.GetAge()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("ageInfo: %v\n", ageInfo)

	var width2, height2 = util.GetImageWidthAndHeight("./hj2.jpg")
	var imageData2 = util.GetResizedBGR("./hj2.jpg")

	// var width2, height2 = util.GetImageWidthAndHeight("./test3.jpg")
	// var imageData2 = util.GetResizedBGR("./test3.jpg")

	faceInfo2, err := engine.DetectFaces(width2-width2%4, height2, ColorFormatBGR24, imageData2)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	singleFaceInfo2 := &SingleFaceInfo{
		FaceRect:   faceInfo2.FaceRect[0],
		FaceOrient: faceInfo2.FaceOrient[0],
	}
	ASFFaceFeature2, faceFeature2, err := engine.ASFFaceFeatureExtract(width2-width2%4, height2, ColorFormatBGR24, imageData2, singleFaceInfo2)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("ASFFaceFeature2: %v\n", ASFFaceFeature2)
	fmt.Printf("faceFeature2: %v\n", faceFeature2)

	// result, err := engine.FaceFeatureCompare(ASFFaceFeature, ASFFaceFeature2)
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return
	// }
	// fmt.Println("result->", result)

	result2, err := engine.FaceFeatureCompare2(ASFFaceFeature, faceFeature2.Feature)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("result2->", result2)

	// 销毁引擎
	if err = engine.Destroy(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
