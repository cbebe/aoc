const std = @import("std");

pub fn main() !void {}

// const data_3a = .{ @embedFile("./input/3/test.txt"), 4361 };
// const data_3b = .{ @embedFile("./input/3/test.txt"), 467835 };
const data_3a = .{ @embedFile("./input/3/input.txt"), 512794 };
const data_3b = .{ @embedFile("./input/3/input.txt"), 67779080 };

const A = @import("aoc.zig");
const split = A.split;

const FoundKey = struct {
    y: i32,
    to: i32,
    from: i32,
};

const GearLocation = struct {
    value: i32,
    y: i32,
    to: i32,
    from: i32,
};

fn solve3a(d: []const u8, alloc: std.mem.Allocator) !i32 {
    var rows = try A.split_to_array(d, "\n", alloc);
    defer rows.deinit();
    for (0.., rows.items) |y, row| {
        std.debug.print("{d}\n", .{y});
        for (0.., row) |x, col| {
            std.debug.print("{c} {d}\n", .{ col, x });
        }
    }

    return 0;
}

fn solve3b(d: []const u8, alloc: std.mem.Allocator) !i32 {
    _ = alloc;
    _ = d;
    return 0;
}

test "day3" {
    try A.testDay(solve3a, data_3a);
    try A.testDay(solve3b, data_3b);
}
