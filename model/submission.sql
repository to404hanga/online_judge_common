CREATE TABLE IF NOT EXISTS submission (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '提交 ID',
    competition_id BIGINT UNSIGNED NOT NULL COMMENT '比赛 ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户 ID',
    problem_id BIGINT UNSIGNED NOT NULL COMMENT '题目 ID',
    code TEXT NOT NULL COMMENT '提交代码',
    stderr TEXT COMMENT '标准错误输出',
    language TINYINT NOT NULL COMMENT '提交语言 ( 0: C, 1: C++, 2: Python, 3: Java, 4: Go )',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '提交状态 ( 0: 待判题, 1: 判题中, 2: 已判题 )',
    result TINYINT NOT NULL DEFAULT 0 COMMENT '判题结果 ( 0: 未判题, 1: Accepted, 2: Wrong Answer, 3: Compile Error, 4: Runtime Error, 5: Time Limit Exceeded, 6: Memory Limit Exceeded, 7: Output Limit Exceeded )',
    time_used INT NOT NULL DEFAULT -1 COMMENT '最大运行时间 ( 单位: 毫秒, -1 表示未判题 )',
    memory_used INT NOT NULL DEFAULT -1 COMMENT '最大运行内存 ( 单位: KB, -1 表示未判题 )',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',

    PRIMARY KEY (id),
    INDEX idx_competition_id_user_id_problem_id (competition_id, user_id, problem_id ASC),
    INDEX idx_created_at (created_at ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='提交表';