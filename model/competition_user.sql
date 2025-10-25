CREATE TABLE IF NOT EXISTS competition_user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛用户 ID',
    competition_id BIGINT UNSIGNED NOT NULL COMMENT '比赛 ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户 ID',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '状态 ( 0: 正常, 1: 禁用 )',
    total_time INT NOT NULL DEFAULT 0 COMMENT '总耗时 ( 单位: 分钟 )',
    pass_count INT NOT NULL DEFAULT 0 COMMENT '通过题目数',

    PRIMARY KEY (id),
    UNIQUE INDEX uk_competition_user_id (competition_id, user_id)
)