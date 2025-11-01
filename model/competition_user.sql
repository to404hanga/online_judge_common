CREATE TABLE IF NOT EXISTS competition_user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛用户 ID',
    competition_id BIGINT UNSIGNED NOT NULL COMMENT '比赛 ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户 ID',
    username varchar(50) NOT NULL COMMENT '用户名(学号)',
    realname varchar(50) NOT NULL COMMENT '真实姓名',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '状态 ( 0: 正常, 1: 禁用 )',
    total_time INT NOT NULL DEFAULT 0 COMMENT '总耗时 ( 单位: 分钟 )',
    pass_count INT NOT NULL DEFAULT 0 COMMENT '通过题目数',

    PRIMARY KEY (id),
    UNIQUE INDEX uk_competition_user_id (competition_id, user_id),
    INDEX idx_ranking (pass_count, total_time) -- 按通过题目数和总耗时排序
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='比赛用户表';