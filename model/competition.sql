CREATE TABLE IF NOT EXISTS competition (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛 ID',
    name VARCHAR(255) NOT NULL COMMENT '比赛名称',
    start_time DATETIME(3) NOT NULL COMMENT '比赛开始时间',
    end_time DATETIME(3) NOT NULL COMMENT '比赛结束时间',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '比赛状态(0:未发布, 1:已发布, 2:已删除)',
    creator_id BIGINT UNSIGNED NOT NULL COMMENT '创建者 ID',
    updater_id BIGINT UNSIGNED NOT NULL COMMENT '更新者 ID',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    PRIMARY KEY (id),
    INDEX idx_start_time (start_time),
    INDEX idx_end_time (end_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='比赛表';
