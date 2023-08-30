[toc]

# 用户表

## tm_user

- u_id `int 11 用户id`
- u_account `varchar 50 账户 唯一`
- u_nick_name `varchar 50 昵称`
- u_password `varchar 200 密码`
- u_status `tinyint 1 1:正常 2:禁用`
- u_avatar `varchar 200 头像地址`
- u_role_type `int 11 1:管理员 2:普通人`
- u_create_time `int 11 创建时间`

### 索引

唯一索引: u_account

# 项目表

## tm_project

- p_id `int 11 项目id`
- p_name `varchar 50 项目名称`
- p_status `tinyint 1 1:正常 2:禁用 3:完结`
- p_create_user_id `int 11 创建的用户id`
- p_create_user_name `varchar 50 创建人昵称`
- p_create_time `int 11 创建时间`
- p_update_time `int 11 更改时间`

# 任务标签表

## tm_task_tag

- tt_id `int 11 任务标签id`
- tt_name `varchar 50 任务标签名字`
- tt_sort `int 11 排序数字`
- tt_color `varchar 20 色号`
- tt_project_id `int 11 所属项目id`

# 任务表

## tm_task

- t_id `int 11 任务id`
- t_name `varchar 50 任务名称`
- t_describe `varchar 200 描述`
- t_sort `int 11 排序数字`
- t_project_id `int 11 项目id`
- t_task_tag_id `int 11 任务标签id`
- t_belong_user_id `int 11 负责人id`
- t_belong_user_name `varchar 50 负责人昵称`
- t_create_user_id `int 11 创建的用户id`
- t_create_user_name `varchar 50 创建人昵称`
- t_status `tinyint 1 1:未完成 2:完成`
- t_create_time `int 11 创建时间`

### 索引

t_project_id

t_task_tag_id

t_belong_user_id

t_create_user_id

# 任务资源表

## tm_task_resource

- tr_id `int 11 资源id`
- tr_absolute_path `varchar 200 绝对路径包含文件名`
- tr_file_name `varchar 20 文件名`
- tr_task_id `int 11 任务id`
- tr_create_time `int 11 创建时间`

### 索引

唯一索引:tr_absolute_path

tr_task_id

# 任务评论表

## tm_task_comment

- tc_id `int 11 评论id`
- tc_task_id `int 11 任务id`
- tc_user_id `int 11 评论人id`
- tc_nick_name `varchar 50 昵称`
- tc_avatar `varchar 200 头像地址`
- tc_comment `varchar 200 评论内容`
- tc_absolute_path `varchar 200 绝对路径包含文件名`
- tc_file_name `varchar 20 文件名`
- tc_create_time `int 11 创建时间`

### 索引

tc_task_id

tc_user_id

tc_resource_id

# 项目人员表

## tm_project_user

- pu_id `int 11 项目人员表id`
- pu_project_id `int 11 项目id`
- pu_user_id `int 11 用户id`

### 索引

pu_project_id
