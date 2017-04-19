"""
Run this script to store whatever is in myfile into db
unforgettables.txt contains a good copy of the unforgetables list
"""

import MySQLdb

conn = MySQLdb.connect(host="192.232.221.54", user="lindjac_lindjac", passwd="0453100594", db="lindjac_dictAPI")
cur = conn.cursor()

ing = ''
file = open("myfile", "r", encoding="utf8")
f = open('myfileoutput.txt', 'w')
direct = 0
for line in file:
    if line[0]=='0':
        img = line[2:-1]

    if line[0] == '1':
        name = line[2:-1]
    if line[0] == '2':
        if direct==1:
            dir = line[2:-1]
            try:
                cur.execute("""INSERT INTO `lindjac_drinkAPI`.`drinks` (`img`, `name`, `time`, `ingredients`, `directions`)
                            VALUES ('""" + img + """', '""" + name + """', '""" + time + """', '""" + ing + """'
                            , '""" + dir + """')""")
                conn.commit()
            except:
                conn.rollback()
            ing = ''
            direct = 0
        else:
            time = line[2:-1]

    if line[0] == '3':
        ing += line[2:-1]
        ing +=","
        direct = 1
file.close()
f.close()
