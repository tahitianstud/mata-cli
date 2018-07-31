package rest_test

//func TestRequestThroughSling(t *testing.T) {
//	Convey("Given a test Request", t, func() {
//
//		testRequest := rest.Request{
//			URL: "https://test.server.com",
//			Payload: "test",
//		}
//
//		Convey("converting should preserve the values", func() {
//			slingRequest := rest.Convert(testRequest)
//
//			So(slingRequest.URL, ShouldEqual, testRequest.URL)
//			So(slingRequest.Payload, ShouldEqual, testRequest.Payload)
//			So(slingRequest.Headers, ShouldEqual, testRequest.Headers)
//			So(slingRequest.Username, ShouldEqual, testRequest.Username)
//			So(slingRequest.Password, ShouldEqual, testRequest.Password)
//		})
//
//		Convey("using a mocked sling client", func() {
//
//			mockedSlinger := mocks.Slinger{}
//			mockedSlinger.On("Get", mock.AnythingOfType("string")).Return(&mockedSlinger)
//			mockedSlinger.On("Post", mock.AnythingOfType("string")).Return(&mockedSlinger)
//			mockedSlinger.On("Set", mock.AnythingOfType("string"), mock.Anything).Return(&mockedSlinger)
//			mockedSlinger.On("BodyJSON", mock.Anything).Return(&mockedSlinger)
//			mockedSlinger.On("ReceiveSuccess", mock.Anything).run(func(args mock.Arguments) {
//				arg := args.Get(0).(*string)
//				*arg = "my value"
//			}).Return(&http.Response{StatusCode:200}, nil)
//
//			mockedSling := &sling.Sling{
//				Client: &mockedSlinger,
//			}
//
//			mockedSlingClient := &rest.SlingClient{
//				Sling: mockedSling,
//			}
//
//			Convey("calling GET with request", func() {
//
//				testTarget := ""
//
//				response, err := mockedSlingClient.Get(testRequest, &testTarget)
//
//				Convey("should return the right values", func() {
//
//					So(response.StatusCode, ShouldEqual, 200)
//					So(err, ShouldBeNil)
//					So(testTarget, ShouldEqual, "my value")
//
//				})
//
//			})
//
//			Convey("calling POST with request", func() {
//
//				testTarget := ""
//
//				response, err := mockedSlingClient.Post(testRequest, &testTarget)
//
//				Convey("should return the right values", func() {
//
//					So(response.StatusCode, ShouldEqual, 200)
//					So(err, ShouldBeNil)
//					So(testTarget, ShouldEqual, "my value")
//
//				})
//
//			})
//
//
//		})
//
//	})
//}
