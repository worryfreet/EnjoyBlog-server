DROP TABLE IF EXISTS `comment`;
CREATE TABLE comment
(
    id                BIGINT       NOT NULL AUTO_INCREMENT COMMENT '自增id',
    comment_id        VARCHAR(50)  NOT NULL UNIQUE COMMENT '评论id',
    comment_content   VARCHAR(255) NOT NULL COMMENT '评论内容',
    commentator_name  VARCHAR(30)  NOT NULL COMMENT '发表评论者',
    commentator_email VARCHAR(40)  NOT NULL COMMENT '发表评论者邮箱',
    support_total     INT          NOT NULL DEFAULT 0 COMMENT '点赞总数',
    article_id        VARCHAR(50)  NOT NULL COMMENT '文章id',
    parent_id         VARCHAR(50)  NOT NULL DEFAULT '0' COMMENT '父评论id, 0默认顶层',
    is_pub            INT          NOT NULL DEFAULT 1 COMMENT '评论是否仅最多三方可看, 1默认公开',
    create_time       DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time       DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_time      DATETIME              DEFAULT NULL,
    PRIMARY KEY (id) USING BTREE,
    KEY (article_id)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;