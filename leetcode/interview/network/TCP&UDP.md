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

### TCP 三次握手

TCP（传输控制协议）三次握手是建立TCP连接的过程，它是确保数据可靠传输的关键步骤。这个过程包括三个主要步骤，分别是SYN、SYN-ACK和ACK。下面将详细解释每一步：

**第一次握手：SYN**

- 发起方（客户端）发送SYN包：握手的第一步是由客户端开始的，客户端选择一个随机的序列号（Seq=X）并向服务器发送一个SYN（同步序列编号）包。这个包的目的是告诉服务器，客户端希望建立连接，并且包含了客户端的初始序列号，这个序列号用于后续的数据传输。
- 目的：主要是为了建立连接，并同步客户端的初始序列号。

**第二次握手：SYN-ACK**
- 接收方（服务器）响应SYN-ACK包：服务器收到客户端的SYN包后，会以自己的SYN包作为响应。服务器的SYN包包含服务器自己的一个随机序列号（Seq=Y），同时对客户端的SYN包进行确认（Ack=X+1，表示服务器期待收到客户端以X+1为序列号的下一个包）。
- 目的：主要是确认客户端的SYN，并同步服务器的序列号。

**第三次握手：ACK**
- 发起方（客户端）发送ACK包：客户端收到服务器的SYN-ACK包后，会发送一个ACK包作为回应。这个ACK包的确认号设置为服务器的SYN包的序列号加一（Ack=Y+1），表示客户端已经成功接收到服务器的SYN包，并期待以Y+1为序列号的服务器的下一个数据包。
- 目的：主要是确认服务器的SYN，并告知服务器客户端准备好接收数据了。

### TCP 四次挥手
**第一次挥手：客户端发送FIN包**
- 发起方：客户端（或称为发起关闭的一方）
- 动作：客户端决定关闭连接，发送一个FIN包给服务器，表示客户端已经没有数据要发送了。
- 状态变化：客户端进入FIN_WAIT_1状态。

**第二次挥手：服务器回应ACK包**
- 发起方：服务器
- 动作：服务器接收到客户端的FIN包后，发送一个ACK包作为响应，表示已经收到客户端的关闭请求。
- 状态变化：客户端收到ACK后，进入FIN_WAIT_2状态。服务器在发送ACK包后，继续发送剩余的数据（如果有的话）。

**第三次挥手：服务器发送FIN包**
- 发起方：服务器
- 动作：服务器发送完剩余的数据后，发送一个FIN包给客户端，表示服务器端也没有数据要发送了。
- 状态变化：服务器发送完FIN包后，进入LAST_ACK状态。

**第四次挥手：客户端回应ACK包**
- 发起方：客户端
- 动作：客户端接收到服务器的FIN包后，发送一个ACK包作为响应。
- 状态变化：客户端发送ACK包后，进入TIME_WAIT状态。在经过一段时间（2倍的MSL，最大报文生存时间）后，确保服务器收到了最后的ACK包，客户端关闭连接，释放资源。

### TCP 常见的11种状态
TCP（传输控制协议）中定义了一系列的状态，以管理TCP连接的生命周期。这些状态涵盖了从连接的建立到连接的终止的整个过程。总共有11种主要的TCP状态：
1. CLOSED：表示没有活动的连接，也没有正在进行的连接。
2. LISTEN：服务器端的套接字处于等待连接请求的状态。
3. SYN_SENT：客户端已发送SYN（同步序列编号）包，正在等待服务器的确认，即已经开始了一个连接请求。
4. SYN_RECEIVED：服务器已收到客户端的SYN包，并发送了SYN+ACK包作为响应，正在等待客户端的确认。
5. ESTABLISHED：表示连接已经建立，数据可以开始传输。
6. FIN_WAIT_1：连接处于正在关闭的过程中，应用程序已经发出了连接终止请求，并发送了FIN包。
7. FIN_WAIT_2：在FIN_WAIT_1状态下，一旦收到对方的ACK包，就进入FIN_WAIT_2状态，此时等待对方的FIN包。
8. CLOSE_WAIT：对方已经初始化了一个连接终止请求，本地端收到了对方的FIN包，并发送了ACK包作为响应，正在等待本地应用程序关闭连接。
9. CLOSING：在同时关闭的情况下，当一方处于FIN_WAIT_1状态时收到了对方的FIN包，此时双方都在尝试关闭连接，会进入CLOSING状态。
10. LAST_ACK：在接收到FIN包并发送ACK包后，等待所有包的确认。
11. TIME_WAIT：表示本地端已发送最后一个ACK包。在这个状态下，会等待足够的时间以确保对方收到了确认包，然后关闭连接。

### TCP 拥塞控制机制
TCP拥塞控制是一种网络拥塞避免机制，确保TCP连接不会过度地向网络发送数据，从而避免网络拥塞。TCP拥塞控制使用一系列算法来动态调整发送方的拥塞窗口大小（Congestion Window, cwnd），以控制在任一时刻向网络发送的数据量。主要的拥塞控制算法包括慢启动（Slow Start）、拥塞避免（Congestion Avoidance）、快速重传（Fast Retransmit）和快速恢复（Fast Recovery）。

**慢启动（Slow Start）**
- 初始阶段：连接开始时，拥塞窗口（cwnd）从1个最大报文段（MSS）开始，每收到一个ACK，cwnd加倍。这导致数据传输速率呈指数增长。
- 阈值：慢启动阈值（ssthresh）是一个上限，当cwnd达到这个阈值时，TCP转换到拥塞避免算法。

**拥塞避免（Congestion Avoidance）**
- 线性增长：当cwnd达到ssthresh后，每收到一个ACK，cwnd增加1/cwnd，导致cwnd呈线性增长，以避免网络过快地进入拥塞状态。
- 拥塞发生：如果发生超时，表明网络拥塞严重，TCP会减少ssthresh的值（通常设置为出现拥塞时的cwnd的一半），cwnd重置为1 MSS，并重新开始慢启动过程。

**快速重传（Fast Retransmit）**
- 重复ACK：当接收方连续收到3个重复的ACK时，表明有一个报文段可能丢失，发送方立即重传该报文段，而不是等待超时发生。
- 不减小cwnd：与超时重传不同，快速重传不会减小cwnd的大小，因为重复ACK表明网络仍然能传输数据包。

**快速恢复（Fast Recovery）**
- 进入快速恢复：在快速重传后，TCP进入快速恢复阶段。此时，ssthresh设置为当前cwnd的一半，cwnd设置为新的ssthresh加上3 MSS（对应于3个重复ACK）。
- 退出快速恢复：如果收到新的ACK，表明丢失的报文段已经被成功接收，TCP退出快速恢复阶段，将cwnd设置为ssthresh的值，进入拥塞避免阶段。

TCP拥塞控制的目标是最大化网络吞吐量，同时避免过度拥塞，确保网络资源的公平使用。通过动态调整发送速率，TCP能够适应网络条件的变化，优化数据传输性能。

### TCP 流量控制机子
流量控制是TCP（传输控制协议）中的一项关键机制，它的主要目的是防止发送方发送数据过快，使得接收方来不及处理，从而避免接收方缓冲区溢出。流量控制确保了发送方和接收方之间的数据传输速率是可控的，以匹配接收方的处理能力。

**实现机制**
- TCP使用滑动窗口协议来实现流量控制。每个TCP连接的接收方都有一个固定大小的缓冲区和一个称为接收窗口（rwnd）的变量，它指示了接收方还能接受多少字节的数据。接收窗口的大小基于接收方的处理速度和缓冲区的剩余空间动态调整，并通过TCP报文段的头部信息通告给发送方。

**工作流程**
- 接收窗口通告：在TCP连接建立时，接收方在SYN报文段中指定初始的接收窗口大小。在后续的通信过程中，每当接收方发送ACK报文段时，它也会在报文段的头部包含当前的接收窗口大小，以通告发送方。
- 发送方控制：发送方根据接收到的最新接收窗口大小调整其发送窗口，确保未被确认的数据量不超过接收方的处理能力。如果接收窗口大小变为0，发送方将停止发送数据，并启动持续计时器（persistence timer）以周期性地探测接收窗口的变化。
- 接收方调整：接收方根据自身缓冲区的使用情况动态调整接收窗口的大小。如果接收方处理数据的速度足够快，它将增加接收窗口的大小；反之，如果缓冲区快要满了，它将减小接收窗口的大小。

**流量控制的挑战**
- 零窗口死锁：如果接收方将窗口大小设置为0，且在一段时间内未能增加窗口大小，发送方将停止发送数据，导致数据传输暂停。为了解决这个问题，TCP实现了持续计时器，发送方会定期发送窗口探测报文段来检查接收窗口是否已增大。
- 窗口更新风暴：当多个接收窗口更新报文段在网络中延迟并且突然到达发送方时，可能导致发送方突然发送大量数据，对网络造成冲击。这种情况较少见，但在网络条件不稳定时可能发生。

流量控制是TCP保证数据可靠传输的重要机制之一，它通过动态调整发送速率来适应接收方的处理能力，确保网络通信的顺畅和高效。