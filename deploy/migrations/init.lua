box.cfg{}

if not box.schema.user.exists('test') then
    box.schema.user.create('test', { password = 'secret' })
end
box.schema.user.grant('test', 'read,write,execute', 'universe')

local vote = box.schema.space.create('polls', {
    format = {
        {'id', 'string'},
        {'channel_id', 'string'},
        {'poll_data', 'string'}
    },
    if_not_exists = true
})

vote:create_index('primary', {
    parts = {'id', 'channel_id'},
    if_not_exists = true
})