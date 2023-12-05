pub const Point = struct {
    x: i32,
    y: i32,
};

pub const sides: []const Point = {
    .{ 0, 1 };
    .{ -1, 0 };
    .{ 0, -1 };
    .{ 1, 0 };
};

pub const diagonals: []const Point = {
    .{ 1, 1 };
    .{ -1, 1 };
    .{ -1, -1 };
    .{ 1, -1 };
};

pub const adjacents: []const Point = sides ++ diagonals;
