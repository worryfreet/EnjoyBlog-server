DROP TABLE IF EXISTS `user`;
CREATE TABLE user (
    id           BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT '自增id',
    user_id      VARCHAR(50)  NOT NULL UNIQUE COMMENT '用户id',
    user_name    VARCHAR(30)  NOT NULL COMMENT '用户姓名',
    email        VARCHAR(40)  NOT NULL UNIQUE DEFAULT '' COMMENT '用户邮箱',
    phone        VARCHAR(20)                  DEFAULT '' COMMENT '手机号',
    password     VARCHAR(100) NOT NULL COMMENT '密码',
    avatar       VARCHAR(100)                 DEFAULT '' COMMENT '用户头像',
    nick_name    VARCHAR(30)                  DEFAULT '' COMMENT '用户昵称',
    description  VARCHAR(255)                 DEFAULT '描述一下优秀的自己吧' COMMENT '用户描述',
    is_admin     INT(1)                       DEFAULT 0 COMMENT '是否是管理员',
    create_time  DATETIME     NOT NULL        DEFAULT CURRENT_TIMESTAMP,
    update_time  DATETIME     NOT NULL        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_time DATETIME                     DEFAULT NULL,
    PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

INSERT INTO user (id, user_id, user_name, email, phone, password)
VALUES (1, '1', '康康', 'worryfreet@163.com', '15539391298', '123456')