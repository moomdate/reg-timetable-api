# API REG SUT STUDENT TIMETABLE 
  API   สำหรับค้นหาตารางเรียนนักศึกษาว่านักศึกษาคนนั้นลงทะเบียนวิชาอะไรบ้าง
  ทำไมถึงทำขึ้นมาในเมื่อก็สามารถเปิดเว็บ http://reg.sut.ac.th เพราะว่าทำให้ง่ายต่อนักพัฒนาที่จะเข้าถึงข้อมูล (เเถ)
## ทำอะไรได้บ้าง 
API จะไปดูข้อมูลของรายวิชาของนักศึกษา เเต่ยังไม่ได้ถึงว่าวิชาที่นักศึกษาคนนั้นลงทะเบียนเรียนตอนไหน
  ```
  https://mycoursetable.herokuapp.com/api/v1/b5917273/2562/1
      {
        _acadyear: "2562",
        _semester: "1",
        course_list: [
            {
                name: "PLURI-CULTURAL THAI STUDIES",
                group: "1",
                course_id: "202324",
                _version: "1"
            },
            {
                name: "PRE-COOPERATIVE EDUCATION",
                group: "1",
                course_id: "523490",
                _version: "1"
            },
            {
                name: "COMPUTER ENGINEERING PROJECT II",
                group: "1",
                course_id: "523496",
                _version: "1"
            }
        ]
    }
  ```

## ใช้งานอย่างไร
    ใช้ HTTP requrest เเบบ GET ที่ URL https://mycoursetable.herokuapp.com/api/v1
    โดยจะรับ Param 3 ตัวคือ 
      1.  รหัสนักศึกษา ต้องขึ้นต้นดัว b B m M 
      2.  ปีการศึกษา
      3.  ภาคการศึกษา
#### ตัวอย่าง
    https://mycoursetable.herokuapp.com/api/v1/b5917273/2562/1

## ที่จะทำเพิ่ม
    - ดึงเวลาของวิชาที่ลงเรียน
    - บา บา บา บา
