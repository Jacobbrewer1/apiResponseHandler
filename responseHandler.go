package apiResponseHandler

func DecodeResponse(code int) response {
	switch code {
	case Response100.Code:
		return Response100
	case Response101.Code:
		return Response101
	case Response102.Code:
		return Response102
	case Response103.Code:
		return Response103
	case Response200.Code:
		return Response200
	case Response201.Code:
		return Response201
	case Response202.Code:
		return Response202
	case Response203.Code:
		return Response203
	case Response204.Code:
		return Response204
	case Response205.Code:
		return Response205
	case Response206.Code:
		return Response206
	case Response207.Code:
		return Response207
	case Response208.Code:
		return Response208
	case Response226.Code:
		return Response226
	case Response300.Code:
		return Response300
	case Response301.Code:
		return Response301
	case Response302.Code:
		return Response302
	case Response303.Code:
		return Response303
	case Response304.Code:
		return Response303
	case Response305.Code:
		return Response305
	case Response306.Code:
		return Response306
	case Response307.Code:
		return Response307
	case Response308.Code:
		return Response308
	case Response400.Code:
		return Response400
	case Response401.Code:
		return Response401
	case Response402.Code:
		return Response402
	case Response403.Code:
		return Response403
	case Response404.Code:
		return Response404
	case Response405.Code:
		return Response405
	case Response406.Code:
		return Response406
	case Response407.Code:
		return Response407
	case Response408.Code:
		return Response408
	case Response409.Code:
		return Response409
	case Response410.Code:
		return Response410
	case Response411.Code:
		return Response411
	case Response412.Code:
		return Response412
	case Response413.Code:
		return Response413
	case Response414.Code:
		return Response414
	case Response415.Code:
		return Response415
	case Response416.Code:
		return Response416
	case Response417.Code:
		return Response417
	case Response418.Code:
		return Response418
	case Response420.Code:
		return Response420
	case Response422.Code:
		return Response422
	case Response423.Code:
		return Response423
	case Response424.Code:
		return Response424
	case Response425.Code:
		return Response425
	case Response426.Code:
		return Response426
	case Response428.Code:
		return Response428
	case Response429.Code:
		return Response429
	case Response431.Code:
		return Response431
	case Response444.Code:
		return Response444
	case Response449.Code:
		return Response449
	case Response450.Code:
		return Response450
	case Response451.Code:
		return Response451
	case Response499.Code:
		return Response499
	case Response500.Code:
		return Response500
	case Response501.Code:
		return Response501
	case Response502.Code:
		return Response501
	case Response503.Code:
		return Response503
	case Response504.Code:
		return Response504
	case Response505.Code:
		return Response505
	case Response506.Code:
		return Response506
	case Response507.Code:
		return Response507
	case Response508.Code:
		return Response508
	case Response510.Code:
		return Response510
	case Response511.Code:
		return Response511
	default:
		UnknownResponse.fmtUnknownResponse(code)
		return UnknownResponse
	}
}
