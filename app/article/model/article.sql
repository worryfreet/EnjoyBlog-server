# 文章表
DROP TABLE IF EXISTS `article`;
CREATE TABLE article
(
    id              BIGINT      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    user_id         VARCHAR(50) NOT NULL UNIQUE COMMENT '用户id',
    article_id      VARCHAR(50) NOT NULL UNIQUE COMMENT '文章id',
    article_content LONGTEXT    NOT NULL COMMENT '文章内容',
    article_title   VARCHAR(50) NOT NULL DEFAULT '' COMMENT '文章标题',
    avatar          VARCHAR(100)         DEFAULT '' COMMENT '文章封面',
    label           VARCHAR(100)         DEFAULT '' COMMENT '文章标签',
    is_top          INT         NOT NULL DEFAULT 0 COMMENT '是否置顶（1置顶，0默认不置顶）',
    is_pub          INT         NOT NULL DEFAULT 0 COMMENT '是否公开（1公开，0默认私密）',
    comment_total   INT         NOT NULL DEFAULT 0 COMMENT '获得评论数量',
    support_total   INT         NOT NULL DEFAULT 0 COMMENT '获得点赞数量',
    create_time     DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time     DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_time    DATETIME             DEFAULT NULL,
    PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

# 文章分类表
DROP TABLE IF EXISTS `article_group`;
CREATE TABLE article_group
(
    id                  BIGINT      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    user_id             VARCHAR(50) NOT NULL UNIQUE COMMENT '用户id',
    article_group_id    VARCHAR(50) NOT NULL UNIQUE COMMENT '分类目录id',
    article_group_title VARCHAR(50) NOT NULL DEFAULT '' COMMENT '分类目录标题',
    parent_id           VARCHAR(50) NOT NULL COMMENT '上级分类目录id',
    create_time         DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time         DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_time        DATETIME             DEFAULT NULL,
    PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

# 文章分类中间表
DROP TABLE IF EXISTS `article_group_rel`;
CREATE TABLE article_group_rel
(
    id               BIGINT(20)  NOT NULL AUTO_INCREMENT COMMENT '自增id',
    article_group_id VARCHAR(50) NOT NULL UNIQUE COMMENT '分类目录id',
    article_id       VARCHAR(50) NOT NULL UNIQUE COMMENT '文章id',
    create_time      DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time      DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_time     DATETIME             DEFAULT NULL,
    PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;