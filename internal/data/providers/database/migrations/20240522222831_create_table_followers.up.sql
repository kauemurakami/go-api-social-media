CREATE TABLE followers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    follower_id UUID,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
            REFERENCES users(id)
            ON DELETE CASCADE, 
    CONSTRAINT fk_follower
        FOREIGN KEY(follower_id) 
            REFERENCES users(id)
            ON DELETE CASCADE 
);