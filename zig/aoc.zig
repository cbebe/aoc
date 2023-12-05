const std = @import("std");

pub const geo = @import("geo.zig");

pub fn splitToArray(x: []const u8, c: []const u8, alloc: std.mem.Allocator) !std.ArrayList([]const u8) {
    var it = std.mem.split(u8, x, c);
    var list = std.ArrayList([]const u8).init(alloc);
    while (it.next()) |elem| {
        if (it.index) |_| {
            try list.append(elem);
        }
    }

    return list;
}

pub fn split(x: []const u8, c: []const u8) std.mem.SplitIterator(u8) {
    return std.mem.split(u8, x, c);
}

pub fn trimChar(x: []const u8, c: []const u8) []const u8 {
    return std.mem.trim(u8, x, c);
}

pub fn trim(x: []const u8) []const u8 {
    return std.mem.trim(u8, x, " ");
}

pub fn int(x: []const u8) !i32 {
    return std.fmt.parseInt(i32, x, 10);
}

const Solver = fn (input: []const u8, alloc: std.mem.Allocator) anyerror!i32;
const TestCase = struct { []const u8, i32 };

pub fn testDay(comptime solver: Solver, comptime data: TestCase) !void {
    try std.testing.expectEqual(@as(i32, data[1]), try solver(data[0], std.testing.allocator));
}
