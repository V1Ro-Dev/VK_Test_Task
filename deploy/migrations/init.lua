local TARANTOOL_USER = os.getenv("TARANTOOL_USER")
local TARANTOOL_PASS = os.getenv("TARANTOOL_PASS")

box.cfg{}

if not box.schema.user.exists(TARANTOOL_USER) then
    box.schema.user.create('test', { password = TARANTOOL_PASS })
end
box.schema.user.grant(TARANTOOL_USER, 'read,write,execute', 'universe')

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