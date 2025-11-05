CREATE TABLE IF NOT EXISTS evaluation (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '判题 ID',
    submission_id BIGINT UNSIGNED NOT NULL COMMENT '提交 ID',
    testcase_number INT NOT NULL COMMENT '测试用例编号',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '评测结果 ( 0: 未评测, 1: Accepted, 2: Wrong Answer, 3: Compile Error, 4: Runtime Error, 5: Time Limit Exceeded, 6: Memory Limit Exceeded, 7: Output Limit Exceeded )',
    time_used INT NOT NULL DEFAULT -1 COMMENT '运行时间 ( 单位: 毫秒, -1 表示未判题 )',
    memory_used INT NOT NULL DEFAULT -1 COMMENT '运行内存 ( 单位: KB, -1 表示未判题 )',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    PRIMARY KEY (id),
    UNIQUE KEY uk_evaluation_submission_testcase (submission_id, testcase_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='判题表';
