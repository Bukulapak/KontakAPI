
# Kontak Fiber (Go) and MongoDB REST API

An API built with Fiber and MongoDB.


## API Usage
Main URL / Endpoint :
https://go-kontak-2997a9d9d62c.herokuapp.com

#### Get all Contact

```http
  GET /kontak
```

#### Get by id Contact

```http
  GET /kontak/{id}
```

#### Add Contact

```http
  POST /insert
```


```json
{
  "nama_kontak": "Si Ujang",
  "nomor_hp": "68123456789",
  "alamat": "Sumedang",
  "keterangan": "Teman"
}

```

#### Edit Contact

```http
  PUT /update/{id}
```


```json
{
  "nama_kontak": "Si Asep",
  "nomor_hp": "61235156431231",
  "alamat": "Bandung",
  "keterangan": "Keluarga"
}

```


#### Delete Contact

```http
  DELETE /delete/{id}
```
