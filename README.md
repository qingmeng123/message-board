# message-board
数据库：message_board中有三个表
user(user_id,username,password,security_answer)
comment(comment_id,post_id,username,txt,comment_time,to_comment_id(回复对象id),state(评论状态：1为默认，存在，0不存在))
post(post_id,username,txt,comment_num,post_time,state)
增加的功能：
通过state来删除留言或评论，评论时对留言的评论数进行更新
入参检验：
通过字典树对回复内容或用户名进行敏感字查询，有密码长度检验
回复套娃：
指定to_comment_id时，对其回复

后续继续加功能。。。
