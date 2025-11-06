CREATE TABLE IF NOT EXISTS competition_problem (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛题目 ID',
    competition_id BIGINT UNSIGNED NOT NULL COMMENT '比赛 ID',
    problem_id BIGINT UNSIGNED NOT NULL COMMENT '题目 ID',
    problem_title  VARCHAR(255) NOT NULL COMMENT '题目标题',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '比赛题目状态 ( 0: 禁用, 1: 启用 )',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    PRIMARY KEY (id),
    UNIQUE uk_competition_problem (competition_id, problem_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='比赛题目表';
