<a name="top"></a>
# QR Service v0.1.0

Generación de QR

- [QR](#qr)
	- [Obtener QR](#obtener-qr)
	


# <a name='qr'></a> QR

## <a name='obtener-qr'></a> Obtener QR
[Back to top](#top)

<p>Obtiene un codigo QR</p>

	GET /qr?code=1234





### Parameter Parameters

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  code | String | <p>Código a generar</p>|
|  size | String | <p>Tamaño 100 200 400 600 800 (opc)</p>|


### Success Response

Respuesta

```
Imagen en formato png
```


### Error Response

500 Server Error

```
HTTP/1.1 500 Internal Server Error
{
   "error" : "Not Found"
}
```
