### Simple Web Crawler

## Inspiration from System Design Interview - Alex Xu


Functional Requirements:
- Given a set of URLs, download (or visit) all web pages addressed by the URLS
- Extract URLS from these web pages
- Add new URLS to the list of URLS to be downloaded
Repeat

Nonfunctional Requirements:
- HTML only
- duplicate URLS should be ignored
- politeness ............... probably not

![System Diagram](diagram.png)

Design Deep Dive:
- BFS since DFS can get too deep
- URL Frontier like a queue with priority and partition tolerance to handle politeness, can expand on this later
- Avoiding duplicate URLS