package icity_sdk

type Header struct {
	Key   string
	Value string
}

var (
	xRequestedWithHeader = Header{
		Key:   "X-Requested-With",
		Value: "XMLHttpRequest",
	}

	acceptHeader = Header{
		Key:   "Accept",
		Value: "application/json, text/javascript, */*; q=0.01",
	}

	iCRenderToSelfHeader = Header{
		Key:   "IC-Render-To",
		Value: "self",
	}
)

func refererHeader(value string) Header {
	return Header{
		Key:   "Referer",
		Value: value,
	}
}

func csrfHeader(value string) Header {
	return Header{
		Key:   "X-CSRF-TOKEN",
		Value: value,
	}
}
