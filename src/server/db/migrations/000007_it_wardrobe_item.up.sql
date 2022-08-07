CREATE TABLE it_wardrobe_item (
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    user_id VARCHAR(50),
    brand_name VARCHAR(255),
    garment_id VARCHAR(50),
    size_type_id VARCHAR(50),
    size_type_item_id VARCHAR(50),
    item_state VARCHAR(255)
);
