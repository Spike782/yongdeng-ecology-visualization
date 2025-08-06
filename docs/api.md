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

  - `risk_level`（筛选风险等级）：支持 `low`, `medium`, `high`，筛选出对应风险等级的区域。

- **示例 URL**:

  - 查询风险等级为 `high` 的所有区域：

    ```bash
    /api/risks?risk_level=high
    ```

- **响应**:

  - **成功**:

```json
{
  "region_id": 1,
  "region_name": "Yongdeng County",
  "risk_index": 0.75,
  "risk_level": "High",
  "data": {
    "ndvi": 0.45,
    "slope": 32.5,
    "land_use": "Agricultural Land"
  }
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

- **URL**: `/api/land-use`

- **方法**: `GET`

- **查询参数**:

  - `region_id`（可选，查询指定区域）：指定区域的唯一标识符。

  - `land_use_type`（筛选土地利用类型）：支持`Vegetation`, `Urban`, `Water`, `Bare Land`，筛选出对应土地利用类型的区域。

- **示例 URL**:

  - 查询特定区域的所有土地利用数据：

    ```bash
    /api/land-use?region_id=1
    ```

  - 查询土地利用类型为 `Agricultural` 的所有区域：

    ```bash
    /api/land-use?land_use_type=Urban
    ```

- **响应**:

  - **成功**:

```json
{
  "region_id": 1,
  "region_name": "Yongdeng County",
  "land_use": "Urban",
  "area": 15000.5,  // 单位：公顷
  "classification": [
    "Crop Land",
    "Grassland",
    "Water Body"
  ]
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
- **请求参数**:

```json
{
  "region_id": 1,
  "report_type": "PDF"  // 可选：PDF 或 Excel
}
```

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