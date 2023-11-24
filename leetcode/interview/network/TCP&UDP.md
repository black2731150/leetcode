# TCP & UDP

### TCP 和 UDP 对比

| 特性             | TCP                                      | UDP                           |
| ---------------- | ---------------------------------------- | ----------------------------- |
| **连接类型**     | 面向连接且可靠                           | 面向非连接且不可靠             |
| **数据传输方式** | 面向字节流                               | 面向报文                       |
| **传输效率**     | 较低（因可靠性导致的额外开销）            | 较高                           |
| **特点**         | 有流量控制、拥塞控制等                    | 无流量控制、无拥塞控制         |
| **应用场景**     | 需要高可靠性但效率要求较低的场景          | 可靠性要求低、效率要求高的场景 |
| **协议应用**     | HTTP、FTP、SMTP                          | DNS、SNMP                     |
