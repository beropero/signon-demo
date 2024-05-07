
DROP DATABASE IF EXISTS signondb;

CREATE DATABASE signondb;

USE signondb;

-- 用户表
CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username   VARCHAR(60) NOT NULL COMMENT "用户名",
    nickname   VARCHAR(60) NOT NULL DEFAULT "未命名用户" COMMENT "昵称",
    password   VARCHAR(60)  COMMENT "密码",
    mobile     VARCHAR(60)  COMMENT "手机号",
    email      VARCHAR(255) COMMENT "邮箱",
    create_at  DATETIME,
    update_at  DATETIME, 
    delete_at  DATETIME
)ENGINE=innoDB DEFAULT CHARSET=utf8;
-- 第三方账号表
/*
    info 第三方账号信息,格式为json字符串
    如微信：
    {
        "openid":"*********"
        "unionid":"*********"
    }
*/
CREATE TABLE third_party_account (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL COMMENT "用户id",
    name VARCHAR(60) NOT NULL COMMENT "平台名",
    account_info JSON NOT NULL COMMENT "第三方账号信息",
    create_at  DATETIME,
    update_at  DATETIME,
    delete_at  DATETIME,
    FOREIGN KEY (user_id) REFERENCES user(id)
)ENGINE=innoDB DEFAULT CHARSET=utf8;
-- jwt令牌存储表
CREATE TABLE jwt_storage (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id   INT NOT NULL COMMENT "用户id",
    token     VARCHAR(1024) NOT NULL COMMENT "jwt令牌",
    logout    BOOLEAN DEFAULT TRUE COMMENT "是否注销",
    expire_at DATETIME NOT NULL COMMENT "失效时间",
    create_at DATETIME,
    update_at DATETIME,
    delete_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES user(id)
)ENGINE=innoDB DEFAULT CHARSET=utf8;
-- 用户登陆日志表
CREATE TABLE user_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id    INT NOT NULL COMMENT "用户id",
    type       VARCHAR(60) NOT NULL COMMENT "登陆类型", 
    ip         VARCHAR(40) NOT NULL COMMENT "ip地址",
    device     VARCHAR(255) COMMENT "设备",
    location   POINT COMMENT "坐标",
    create_at  DATETIME,
    update_at  DATETIME,
    delete_at  DATETIME,
    FOREIGN KEY (user_id) REFERENCES user(id)
)ENGINE=innoDB DEFAULT CHARSET=utf8;
-- 验证码请求记录表
CREATE TABLE code_record (
    id INT AUTO_INCREMENT PRIMARY KEY,
    account     VARCHAR(60)  COMMENT "账号",
    code       VARCHAR(10)  COMMENT "验证码",
    create_at  DATETIME,
    update_at  DATETIME,
    delete_at  DATETIME
)ENGINE=innoDB DEFAULT CHARSET=utf8;

INSERT INTO user(username,nickname,password,mobile,email) VALUES
            ("zhangsan","xiaoming","12345678","13542107800","1578999999@qq.com");

INSERT INTO third_party_account(user_id,name,account_info) VALUES  
            ("1","Wechat",'{"openid":"*********","unionid":"*******"}'),
            ("1","Wechat",'{"openid":"*********","unionid":"*******"}'),
            ("1","Wechat",'{"openid":"*********","unionid":"*******"}');