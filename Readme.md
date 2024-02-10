# OTP Service API

Bu proje, OTP (One Time Password) hizmetini sunan bir API'yi içerir. Twilio kullanarak SMS gönderme ve doğrulama işlevselliğini sağlar.

## Kullanım

1. `.env` Dosyası: Projenin çalışması için gerekli olan çevresel değişkenleri içerir. Aşağıdaki değişkenler bu dosyada tanımlanmalıdır:

    ```plaintext
    ACCOUNT_SID= "Twilio hesap ID'nizi buraya yazın"
    AUTHTOKEN = "Twilio oturum açma belirtecinizi buraya yazın"
    SERVICESID="Twilio doğrulama hizmeti ID'nizi buraya yazın"
    ```

    Lütfen bu değerleri kendi Twilio hesabınızın bilgileriyle değiştirin.

2. `.env` Dosyasının yüklenmesi: `.env` dosyasındaki değişkenlerin yüklenmesi için `envAccount` adlı bir fonksiyon bulunmaktadır. Bu fonksiyon `.env` dosyasını yükler ve belirtilen değişken adını alır.

## API Paketi Yapısı

Proje, aşağıdaki yapıya sahip olan `api` paketi içerir:

- `Config`: API yapılandırmasını tanımlar ve `sendSMS` ve `verifySMS` gibi HTTP isteklerine yanıt vermek için yönlendiricileri yönetir.
- `LoggingService`: Twilio API çağrılarını günlüklemek için bir servis sağlar.
- `SMSService`: Twilio ile SMS gönderme ve doğrulama işlevselliğini uygular.
- `jsonResponse`: JSON yanıtları için yapısal bir şablon sağlar.
- `validateBody`, `writeJSON`, `errorJSON`: HTTP isteklerini doğrular ve yanıtları oluşturmak için yardımcı işlevler sağlar.

## SMS Gönderme Endpoint'i (/otp)


###    URL: POST /otp


### Request Body:
    
```bash 
    {
    "phoneNumber": "+905551234567"
    }




###  Response (HTTP 202 Accepted):
        ```bash
        {
            "status": 202,
        "message": "success",
        "data": "OTP verified successfully"
         }

## Başlangıç

Proje, `main.go` dosyasında ana uygulamayı başlatır. `gin` yönlendiricisi üzerinde bir HTTP sunucusu başlatır ve API'yi dinlemeye başlar.

## Lisans

Bu proje MIT Lisansı altında lisanslanmıştır. Detaylar için [LICENSE](LICENSE) dosyasına bakınız.
