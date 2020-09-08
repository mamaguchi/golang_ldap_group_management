package main

import (
        "log"
        "fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/google/uuid"
)

func genUUID() string {
    uid := uuid.New()
    return uid.String()
}

func addDynGrp() {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    err = l.Bind("cn=Directory Manager", "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    uid := genUUID()
    newDn := fmt.Sprintf("uid=%s,ou=groups,dc=example,dc=com", uid)

    addRequest := ldap.NewAddRequest(
	    newDn,
	    nil,
    )
    addRequest.Attribute("objectClass", []string{"top"})
    addRequest.Attribute("objectClass", []string{"groupOfURLs"})
    addRequest.Attribute("objectClass", []string{"uidObject"})
    addRequest.Attribute("uid", []string{uid})
    addRequest.Attribute("cn", []string{"dynamicGroup"})
    addRequest.Attribute("ou", []string{"groups"})

    grpOwnerDn := "staffId=880601105149,ou=People,dc=example,dc=com"
    addRequest.Attribute("owner", []string{grpOwnerDn})

    memberURL := fmt.Sprintf("ldap:///ou=People,dc=example,dc=com??sub?groupUid=%s", uid)
    addRequest.Attribute("memberURL", []string{memberURL})

    err = l.Add(addRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func addDefaultUser() {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    err = l.Bind("cn=Directory Manager", "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    defaultUserDn := "staffId=880601105149,ou=People,dc=example,dc=com"

    addRequest := ldap.NewAddRequest(
	    defaultUserDn,
	    nil,
    )
    addRequest.Attribute("objectClass", []string{"top"})
    addRequest.Attribute("objectClass", []string{"staff"})
    addRequest.Attribute("staffName", []string{"patrickchow"})
    addRequest.Attribute("staffId", []string{"880601105149"})
    addRequest.Attribute("staffTel", []string{"0163184120"})

    err = l.Add(addRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func addMemberUser() {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    err = l.Bind("cn=Directory Manager", "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    memberUserDn := "staffId=880601105150,ou=People,dc=example,dc=com"

    addRequest := ldap.NewAddRequest(
	    memberUserDn,
	    nil,
    )
    addRequest.Attribute("objectClass", []string{"top"})
    addRequest.Attribute("objectClass", []string{"staff"})
    addRequest.Attribute("staffName", []string{"mamaguchi"})
    addRequest.Attribute("staffId", []string{"880601105150"})
    addRequest.Attribute("staffTel", []string{"911"})

    err = l.Add(addRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func addMemberUserToDynGrp() {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    err = l.Bind("cn=Directory Manager", "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    memberUserDn := "staffId=880601105150,ou=People,dc=example,dc=com"
    modifyRequest := ldap.NewModifyRequest(memberUserDn, nil)

    dynamicGroupUid := "b50ff8cc-3572-4583-a2be-c9e3252bc2dd"
    modifyRequest.Add("groupUid", []string{dynamicGroupUid})

    err = l.Modify(modifyRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func addMemberToDynGrpByRootAdm(staffId string, groupUid string) {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    err = l.Bind("cn=Directory Manager", "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    memberDN := fmt.Sprintf("staffId=%s,ou=People,dc=example,dc=com", staffId)
    modifyRequest := ldap.NewModifyRequest(memberDN, nil)

    modifyRequest.Add("groupUid", []string{groupUid})

    err = l.Modify(modifyRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func addMemberToDynGrpByGrpOwner(staffId string, groupUid string, grpOwnerId string) {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    grpOwnerDN := fmt.Sprintf("staffId=%s,ou=People,dc=example,dc=com", grpOwnerId)
    err = l.Bind(grpOwnerDN, "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    memberDN := fmt.Sprintf("staffId=%s,ou=People,dc=example,dc=com", staffId)
    modifyRequest := ldap.NewModifyRequest(memberDN, nil)

    modifyRequest.Add("groupUid", []string{groupUid})

    err = l.Modify(modifyRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func deleteMemberFromDynGrpByGrpOwner(staffId string, groupUid string, grpOwnerId string) {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    grpOwnerDN := fmt.Sprintf("staffId=%s,ou=People,dc=example,dc=com", grpOwnerId)
    err = l.Bind(grpOwnerDN, "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    memberDN := fmt.Sprintf("staffId=%s,ou=People,dc=example,dc=com", staffId)
    modifyRequest := ldap.NewModifyRequest(memberDN, nil)

    modifyRequest.Delete("groupUid", []string{groupUid})

    err = l.Modify(modifyRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func addGrpToMytcaGroup(newGrpCN string) {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    grpOwnerDn := "staffId=880601105149,ou=People,dc=example,dc=com"
    err = l.Bind(grpOwnerDn, "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    uid := genUUID()
    newGrpDN := fmt.Sprintf("uid=%s,ou=mytcaGroup,ou=groups,dc=example,dc=com", uid)

    addRequest := ldap.NewAddRequest(
	    newGrpDN,
	    nil,
    )
    addRequest.Attribute("objectClass", []string{"top"})
    addRequest.Attribute("objectClass", []string{"groupOfURLs"})
    addRequest.Attribute("objectClass", []string{"uidObject"})
    addRequest.Attribute("uid", []string{uid})
    addRequest.Attribute("cn", []string{newGrpCN})
    addRequest.Attribute("ou", []string{"mytcaGroup"})

    addRequest.Attribute("owner", []string{grpOwnerDn})

    memberURL := fmt.Sprintf("ldap:///ou=People,dc=example,dc=com??sub?groupUid=%s", uid)
    addRequest.Attribute("memberURL", []string{memberURL})

    err = l.Add(addRequest)
    if err != nil {
	    log.Fatal(err)
    }
}

func delGrpFromMytcaGroup(grpUID string, grpOwnerId string) {
    l, err := ldap.DialURL("ldap://127.0.0.1:389")
    if err != nil {
	    log.Fatal(err)
    }
    defer l.Close()

    grpOwnerDn := fmt.Sprintf("staffId=%s,ou=People,dc=example,dc=com", grpOwnerId)
    err = l.Bind(grpOwnerDn, "88motherfaker88")
    if err != nil {
	    log.Fatal(err)
    }

    grpDN := fmt.Sprintf("uid=%s,ou=mytcaGroup,ou=groups,dc=example,dc=com", grpUID)
    del := ldap.NewDelRequest(grpDN, nil)

    err = l.Del(del)
    if err != nil {
            log.Fatal(err)
    }
}

func main() {
	//uid := genUUID()
        //fmt.Printf("Generated UUID: %s\n", uid)
	//addDefaultUser()
	//addDynGrp()
	//addMemberUser()
        //addMemberUserToDynGrp()
	//addMemberToDynGrpByRootAdm("880601105149", "masterAdm")
	//addGrpToMytcaGroup("mytcaSubgroup1")
        //delGrpFromMytcaGroup("d742ee02-03a2-4b54-a31b-86dbc897d06c", "880601105149")
        //addMemberToDynGrpByGrpOwner("880601105150", "db1986de-3e16-41de-aea3-c67c01a5b828", "880601105149")
        deleteMemberFromDynGrpByGrpOwner("880601105150", "db1986de-3e16-41de-aea3-c67c01a5b828", "880601105149")
}


