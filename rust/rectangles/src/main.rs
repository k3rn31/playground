#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn square(size: u32) -> Rectangle {
        return Rectangle {
            width: size,
            height: size,
        };
    }

    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn perimeter(&self) -> u32 {
        self.width * 2 + self.height * 2
    }

    fn can_hold(&self, rectangle: &Rectangle) -> bool {
        return rectangle.width <= self.width && rectangle.height <= self.height;
    }
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    let rect2 = Rectangle {
        width: 10,
        height: 40,
    };
    let rect3 = Rectangle {
        width: 60,
        height: 45,
    };
    let sq = Rectangle::square(10);

    println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
    println!("Can rect1 hold rect3? {}", rect1.can_hold(&rect3));
    println!("Can rect1 hold sq? {}", rect1.can_hold(&sq));

    println!("rect1 area is {}", rect1.area());
    println!("rect1 perimeter is {}", rect1.perimeter());
}
