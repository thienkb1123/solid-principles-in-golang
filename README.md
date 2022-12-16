# solid-principles-in-golang

1. Single Responsibility Principle (SRP) - Nguyên lý Đơn Trách Nhiệm
    > **_NOTE:_**  Nguyên lý này ứng với chữ S trong SOLID. Có ý nghĩa là một class chỉ nên giữ một trách nhiệm (chức năng) duy nhất. Mỗi lớp chỉ nên chịu trách nhiệm về một nhiệm vụ cụ thể nào đó. Đề xác định một class giữ một hay nhiều trách nhiệm bằng cách xem xét nếu có thể nghĩ ra nhiều hơn một động cơ làm thay đổi class, thì class đó có nhiều hơn một trách nhiệm.

2. Open/Closed Principle (OCP) -  Nguyên lý Đóng Mở
    > **_NOTE:_** Nguyên lý thứ hai, tương ứng với chữ O trong SOLID. Chúng ta nên hạn chế việc chỉnh sửa bên trong một class hoặc module có sẵn, thay vào đó chúng ta hay xem xét việc mở rộng chúng. Hạn chế sửa đổi: Không nên chỉnh sửa source code của một module hoặc một class có sẵn, vì sẽ làm ảnh hướng đến tính đúng đắn của chương trình. Ưu tiên mở rộng: Khi cần thêm tính năng mới, nên kế thừa và mở rộng các module/class có sẵn thành các module con lớn hơn. Các module/class con vừa có các đặc tính của lớp cha (đã được kiểm chứng đúng đắn), vừa được bổ sung tính năng mới phù hợp với yêu cầu.


3. Liskov Substitution Principle (LSP) - Nguyên lý Thay Thế Lít Kốp 
    > **_NOTE:_** Nguyên lý thứ ba, tương ứng với chữ L trong SOLID. Trong một chương trình, các object của class con có thể thay thế class cha mà không làm thay đổi tính đúng đắn của chương trình. Để giữ tính đúng đắn của chương trình, class con phải thay thế được class cha.

4. Interface Segregation Principle (ISP) - Nguyên lý phân tách interface.
    > **_NOTE:_**  Nguyên lý thứ tư, tương ứng với chữ I trong SOLID. Thay vì dùng 1 interface lớn, ta nên tách thành nhiều interface nhỏ, với nhiều mục đích cụ thể. Interface là một lớp rỗng chỉ chứa khai báo về tên phương thức không có khai báo về thuộc tính hay thứ gì khác và các phương thức này cũng là rỗng. Bởi vậy bất kỳ lớp nào sử dụng lớp interface đều phải định nghĩa các phương thức đã khai báo ở lớp interface. Để thiết kế một hệ thống linh hoạt, dễ thay đổi, các module của hệ thống nên giao tiếp với nhau thông qua interface. Mỗi module sẽ gọi chức năng của module khác thông qua interface mà không cần quan tâm tới implementation bên dưới.

5. Dependency Inversion Principle (DIP) - Nguyên lý Đảo Ngược Dependency
    > **_NOTE:_**  Nguyên lý cuối cùng, tương ứng với chữ D trong SOLID. Các module cấp cao không nên phụ thuộc vào các module cấp thấp. Cả 2 nên phụ thuộc vào abstraction. Interface (abstraction) không nên phụ thuộc vào chi tiết, mà ngược lại. (Các class giao tiếp với nhau thông qua interface, không phải thông qua implementation.)