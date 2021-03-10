<a name="top"></a>
# QR Service v0.1.0

Generación de QR

- [Barcode](#barcode)
	- [Obtener Barcode](#obtener-barcode)
	


# <a name='barcode'></a> Barcode

## <a name='obtener-barcode'></a> Obtener Barcode
[Back to top](#top)

<p>Obtiene un barcode</p>

	GET /barcode/:barcode



### Examples

size : Parametro url o header

```
size=[100|200|400|600|800]
```


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
