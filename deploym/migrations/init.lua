box.cfg{}

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