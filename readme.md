Image upload with Resize
===

### 環境建置
* MacOS Mojave: 10.14
* Golang: 1.11.5
* Framework: Gin
* Docker: 18.09.1

### 啟動環境
```dockerfile=
cd uploader

//build image
docker build . -t sam774410/uploader

//將上傳圖片及任意尺寸圖片 掛載到桌面的test資料夾
docker run -d -p 3000:3000 -v ~/Desktop/test:/app/upload/orign -v ~/Desktop/test:/app/upload/custom sam774410/uploader

```
### 上傳圖片
* 進入連結：127.0.0.1:3000 即可上傳圖片

* API
    * Method ：POST
    * URL：[POST] 127.0.0.1:3000/upload
    * Content-Type：multipart/form-data
    * key=**upload**; file=**your image**

### 更改圖片大小
* 圖片掛載到test資料夾，filename = [Date]_filename
    * Example：20190413_car.png 
* 原圖：
    * 進入連結：127.0.0.1:3000/image/filename
    ```
    Example:
    127.0.0.1:3000/image/20190412_car.png
    ```
    
    
* 任意大小圖：
    * 進入連結：127.0.0.1:3000/image/filename?width=任意寬度&height=任意高度
    ```
    Example:
    127.0.0.1:3000/image/20190412_car.png?width=500&height=800
    ```