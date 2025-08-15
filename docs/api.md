# API 文档

## 1. 用户认证模块

### 1.1 用户注册

- **URL**: `/api/auth/register`
- **方法**: `POST`
- **请求参数**:

```json
{
  "username": "xiaoming",
  "email": "xiaoming@example.com",
  "password": "password123"
}
```

- **响应**:
  - **成功**:

```json
{
  "message": "Registration successful",
  "user": {
    "username": "xiaoming",
    "email": "xiaoming@example.com"
  }
}
```

- **失败**:

```json
{
  "error": "Email already exists"
}
```

### 1.2 用户登录

- **URL**: `/api/auth/login`
- **方法**: `POST`
- **请求参数**:

```json
{
  "username": "xiaoming",
  "password": "password123"
}
```

- **响应**:
  - **成功**:

```json
{
  "message": "Login successful",
  "token": "JWT_TOKEN"
}
```

- **失败**:

```json
{
  "error": "Invalid username or password"
}
```

## 2. 生态风险评估模块

### 2.1 获取区域的生态风险数据（支持筛选风险等级）

- **URL**: `/api/risks`

- **方法**: `GET`

- **查询参数**:

  - `gridcode`（筛选风险等级）：`gridcode`=1,2,3,4，对应 极低风险，低风险，中风险，高风险地区，筛选出对应风险等级的区域。

- **示例 URL**:

  - 查询风险等级为高的所有区域：

    ```bash
    /api/risks/gridcode?gridcode=4
    ```

- **响应**:

  - **成功**:

```json
{
    "code": 200,
    "count": 25463,
    "data": [
        {
            "ID": 8,
            "Geom": "MULTIPOLYGON(((103.40394153915992 36.99618567059446,103.40394153915992 36.99679925294578,103.40472739194706 36.99649246177,103.40394153915992 36.99618567059446)))",
            "Id": 8,
            "Gridcode": 2,
            "ShapeLeng": 0.00230081136,
            "ShapeArea": 2.4109e-7
        },
        {
            "ID": 35,
            "Geom": "MULTIPOLYGON(((103.38307973921195 36.991890594134475,103.38233163612051 36.992100641182844,103.38286971908087 36.99263868844878,103.38307973921195 36.991890594134475)))",
            "Id": 35,
            "Gridcode": 2,
            "ShapeLeng": 0.00231498631,
            "ShapeArea": 2.5777e-7
        },
        ......
      ],
        "message": "查询成功"
}

```

- **失败**:

```json
{
  "error": "Region not found"
}
```

## 3. 土地利用模块

### 3.1 获取区域土地利用数据（支持筛选土地利用类型）

- **URL**: `/api/usages`

- **方法**: `GET`

- **查询参数**:

  - `gridcode`（筛选土地利用类型）：`gridcode`=1,2,3,4，对应植被，水体，裸土，人工建筑地区，筛选出对应土地利用类型的区域。

- **示例 URL**:

  - 查询土地利用类型为人工建筑的所有区域：

    ```bash
    /api/usages/gridcode?gridcode=4
    ```

- **响应**:

  - **成功**:

```json
{
    "code": 200,
    "count": 59706,
    "data": [
        {
            "ID": 125088,
            "Geom": "MULTIPOLYGON(((103.66630694636933 36.405225678042314,103.66555884327789 36.40543573562337,103.66609692623848 36.40597378113392,103.66630694636933 36.405225678042314)))",
            "Id": 125088,
            "Gridcode": 4,
            "ShapeLeng": 0.00231499637,
            "ShapeArea": 2.5777e-7
        },
        {
            "ID": 125091,
            "Geom": "MULTIPOLYGON(((103.67428351693786 36.407066425096616,103.67509843099833 36.406654174454275,103.67551068164067 36.405839260393805,103.67450155138914 36.40533761637539,103.67428351693786 36.407066425096616)))",
            "Id": 125091,
            "Gridcode": 4,
            "ShapeLeng": 0.00469595304,
            "ShapeArea": 0.00000117405
        },
        ......
        ],
        "message": "查询成功"
}
```

- **失败**:

```json
{
  "error": "Region not found"
}
```

## 4. 生态报告模块

### 4.1 生成生态风险评估报告

- **URL**: `/api/reports/eco-risk`
- **方法**: `POST`
- **请求参数**:"filename"

- **响应**:
  - **成功**:

```json
{
  "message": "Report generated successfully",
  "download_url": "https://example.com/reports/eco-risk-report-1.pdf"//下载链接
}
```

- **失败**:

```json
{
  "error": "Failed to generate report"
}
```

------

## 错误处理

所有 API 错误响应遵循以下格式：

- **400 Bad Request**：请求参数无效或缺失。
- **401 Unauthorized**：未授权，用户未登录或权限不足。
- **404 Not Found**：请求的资源不存在。
- **500 Internal Server Error**：服务器内部错误。