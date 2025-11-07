// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-05, by liasica

package gitlab

type MergeState = string

const (
	MergeStateOpened MergeState = "opened"
	MergeStateClosed MergeState = "closed"
	MergeStateLocked MergeState = "locked"
	MergeStateMerged MergeState = "merged"
)

// DetailedMergeStatus 表示 Merge Request 的详细合并状态（GitLab 15.6+ 推荐使用）
// detailed_merge_status
type DetailedMergeStatus = string

const (
	DetailedMergeStatusApprovalsSyncing         DetailedMergeStatus = "approvals_syncing"          // 合并请求的审批正在同步中
	DetailedMergeStatusChecking                 DetailedMergeStatus = "checking"                   // Git 正在检查是否可以进行有效的合并
	DetailedMergeStatusCIMustPass               DetailedMergeStatus = "ci_must_pass"               // 合并前必须通过 CI/CD 流水线
	DetailedMergeStatusCIStillRunning           DetailedMergeStatus = "ci_still_running"           // CI/CD 流水线仍在运行中
	DetailedMergeStatusCommitsStatus            DetailedMergeStatus = "commits_status"             // 源分支应存在并包含提交
	DetailedMergeStatusConflict                 DetailedMergeStatus = "conflict"                   // 源分支与目标分支存在冲突
	DetailedMergeStatusDiscussionsNotResolved   DetailedMergeStatus = "discussions_not_resolved"   // 所有讨论未解决，无法合并
	DetailedMergeStatusDraftStatus              DetailedMergeStatus = "draft_status"               // 当前为草稿状态，不能合并
	DetailedMergeStatusJiraAssociationMissing   DetailedMergeStatus = "jira_association_missing"   // 标题或描述必须引用 Jira issue
	DetailedMergeStatusMergeable                DetailedMergeStatus = "mergeable"                  // 分支可正常合并
	DetailedMergeStatusMergeRequestBlocked      DetailedMergeStatus = "merge_request_blocked"      // 被其他合并请求阻塞
	DetailedMergeStatusMergeTime                DetailedMergeStatus = "merge_time"                 // 未到可合并的时间点
	DetailedMergeStatusNeedRebase               DetailedMergeStatus = "need_rebase"                // 合并请求必须先进行 rebase
	DetailedMergeStatusNotApproved              DetailedMergeStatus = "not_approved"               // 尚未获得审批，无法合并
	DetailedMergeStatusNotOpen                  DetailedMergeStatus = "not_open"                   // 合并请求未处于打开状态
	DetailedMergeStatusPreparing                DetailedMergeStatus = "preparing"                  // 正在准备合并请求差异数据
	DetailedMergeStatusRequestedChanges         DetailedMergeStatus = "requested_changes"          // 审阅者已请求更改
	DetailedMergeStatusSecurityPolicyViolations DetailedMergeStatus = "security_policy_violations" // 存在安全策略违规
	DetailedMergeStatusStatusChecksMustPass     DetailedMergeStatus = "status_checks_must_pass"    // 所有状态检查必须通过才能合并
	DetailedMergeStatusUnchecked                DetailedMergeStatus = "unchecked"                  // Git 尚未检查是否可合并
	DetailedMergeStatusLockedPaths              DetailedMergeStatus = "locked_paths"               // 目标分支中的锁定路径需解锁
	DetailedMergeStatusLockedLFSFiles           DetailedMergeStatus = "locked_lfs_files"           // 被其他用户锁定的 LFS 文件需解锁
	DetailedMergeStatusTitleRegex               DetailedMergeStatus = "title_regex"                // 标题未匹配项目配置的正则规则
)
