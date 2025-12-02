CREATE TABLE IF NOT EXISTS competition_user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛用户 ID',
    competition_id BIGINT UNSIGNED NOT NULL COMMENT '比赛 ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户 ID',
    username varchar(50) NOT NULL COMMENT '用户名(学号)',
    realname varchar(50) NOT NULL COMMENT '真实姓名',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '状态 ( 0: 正常, 1: 禁用 )',
    pass_count INT NOT NULL DEFAULT 0 COMMENT '通过题目数',
    total_time BIGINT NOT NULL DEFAULT 0 COMMENT '总耗时 ( 单位: 毫秒 )',
    retry_count INT NOT NULL DEFAULT 0 COMMENT '重试次数',
    start_time DATETIME NOT NULL COMMENT '比赛开始时间',

    PRIMARY KEY (id),
    UNIQUE INDEX uk_competition_user_id (competition_id, user_id),
    INDEX idx_ranking (competition_id ASC, pass_count DESC, total_time ASC) -- 按通过题目数和总耗时排序
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='比赛用户表';