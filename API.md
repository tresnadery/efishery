# API Documentation

## Auth Service

### Register [/users/registraions]

#### Create New User[POST]

* Role ID
```
	Admin : 74fd1f8a-c8ea-4728-ad8a-147ea82a5c30
	Customer : 40d6568b-410f-4093-a75c-52671bcc9648
```

Add new user if name is not exist

* Request
```
{
	"phone_number":"081234567891",
	"name":"admin1",
	"role_id":"74fd1f8a-c8ea-4728-ad8a-147ea82a5c30"
}
```

* Response 200 (application/json)
```
{
	"message": "User Successfully Created",
	"password" : "rJ8S"
}
```

### Access Token [/tokens]

#### Generate Access Token[POST]

Get access token use phone number and password

* Request
```
	{
		"phone_number":"081234567891",
    	"password": "rJ8S"
	}
```

* Response 200 (application/json)
```
	{
		"message": "Access Token Successfully Created",
    	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwODUxMTExMTExMTExMTEiLCJuYW1lIjoiYWRtaW4xIiwicm9sZV9uYW1lIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiMjAyMC0wNy0yMVQwMjowMjo1MS42NjU1OVoiLCJleHAiOjE1OTg4OTY5ODZ9.2RfEirS-V5theutxNuppWF_hZCu19oqQg_BIUq7JzZk",
    	"token_type": "Bearer"
	}
```

### Profile [/users/profiles]

#### Get Profiles From Access Token[GET]

* Request

* Headers

```
	Authorization : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwODUxMTExMTExMTExMTEiLCJuYW1lIjoiYWRtaW4xIiwicm9sZV9uYW1lIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiMjAyMC0wNy0yMVQwMjowMjo1MS42NjU1OVoiLCJleHAiOjE1OTg4OTY5ODZ9.2RfEirS-V5theutxNuppWF_hZCu19oqQg_BIUq7JzZk
```

* Response 200 (application/json)

```
	{
		"message": "Get profile successfully",
    	"user": {
        	"created_at": "2020-07-21T02:02:51.66559Z",
        	"name": "admin1",
        	"phone_number": "085111111111111",
        	"role_name": "admin"
    	}
	}
```

## Storage Service

### Storage [/storages]

#### Get Storage[GET]

* Request

* Headers

```
	Authorization : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwODUxMTExMTExMTExMTEiLCJuYW1lIjoiYWRtaW4xIiwicm9sZV9uYW1lIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiMjAyMC0wNy0yMVQwMjowMjo1MS42NjU1OVoiLCJleHAiOjE1OTg4OTY5ODZ9.2RfEirS-V5theutxNuppWF_hZCu19oqQg_BIUq7JzZk
```

* Response 200 (application/json)

```
	{
		"message": "Get storage successfully",
		"data": [
		    {
		        "uuid": "",
		        "komoditas": "",
		        "area_provinsi": "",
		        "area_kota": "",
		        "size": "",
		        "price": "20000",
		        "price_in_usd": "1.35",
		        "tgl_parsed": "",
		        "timestamp": ""
		    },
        	{
	            "uuid": "8a23fcab-ef67-48b8-8ba1-7055ea91ea3b",
	            "komoditas": "Ikan Tunaa",
	            "area_provinsi": "JAWA TIMUR",
	            "area_kota": "SURABAYA",
	            "size": "90",
	            "price": "20000",
	            "price_in_usd": "1.35",
	            "tgl_parsed": "Wed Jun 03 11:32:48 GMT+07:00 2020",
	            "timestamp": "1591158768"
        	},
        	....
        ]
	}
```

### Storage With Aggregate [/admin/storages]

#### Get Storage[GET]

* Request

* Headers

```
	Authorization : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwODUxMTExMTExMTExMTEiLCJuYW1lIjoiYWRtaW4xIiwicm9sZV9uYW1lIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiMjAyMC0wNy0yMVQwMjowMjo1MS42NjU1OVoiLCJleHAiOjE1OTg4OTY5ODZ9.2RfEirS-V5theutxNuppWF_hZCu19oqQg_BIUq7JzZk
```

* Response 200 (application/json)

```
	{
		"message": "Get storage successfully",
		"data": [
		     {
	            "province_area": "JAWA BARAT",
	            "year": "52280",
	            "week": "14",
	            "min": "20000",
	            "max": "20000",
	            "median": "1.000000",
	            "avg": "20000.000000"
	        },
	        {
	            "province_area": "JAWA BARAT",
	            "year": "52304",
	            "week": "44",
	            "min": "20000",
	            "max": "20000",
	            "median": "1.000000",
	            "avg": "20000.000000"
	        },
        	....
        ]
	}
```

### Profile

#### Get Profiles From Access Token[GET]

* Request

* Headers

```
	Authorization : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwODUxMTExMTExMTExMTEiLCJuYW1lIjoiYWRtaW4xIiwicm9sZV9uYW1lIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiMjAyMC0wNy0yMVQwMjowMjo1MS42NjU1OVoiLCJleHAiOjE1OTg4OTY5ODZ9.2RfEirS-V5theutxNuppWF_hZCu19oqQg_BIUq7JzZk
```

* Response 200 (application/json)

```
	{
		"message": "Get profile successfully",
    	"user": {
        	"created_at": "2020-07-21T02:02:51.66559Z",
        	"name": "admin1",
        	"phone_number": "085111111111111",
        	"role_name": "admin"
    	}
	}
```
