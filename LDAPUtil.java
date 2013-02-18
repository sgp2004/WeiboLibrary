import java.util.Hashtable;
import javax.naming.Context;
import javax.naming.NamingEnumeration;
import javax.naming.directory.Attribute;
import javax.naming.directory.Attributes;
import javax.naming.directory.DirContext;
import javax.naming.directory.InitialDirContext;


/**
 *
 * @author x-spirit
 */
public class LDAPUtil {

    private Hashtable<String,String> env = null;
    DirContext ctx = null;
    boolean bLogin = false;
    boolean getAttr = false;

    public LDAPUtil(String name, String password) {
	if (name != null && name.contains("@staff.sina")){
	    name = name.substring(0, name.indexOf('@'));
	}
        env = new Hashtable<String,String>();
        env.put(Context.INITIAL_CONTEXT_FACTORY, "com.sun.jndi.ldap.LdapCtxFactory");
        env.put(Context.PROVIDER_URL, "ldap://10.210.96.10");
        env.put(Context.SECURITY_AUTHENTICATION, "Simple");
        env.put(Context.SECURITY_PRINCIPAL, "uid=" + name + ",ou=people,o=staff.sina.com.cn,o=usergroup");
        env.put(Context.SECURITY_CREDENTIALS, password);
    }//end ADAuth()

    public LDAPUtil() {
        this("Administrator", "xxxxxxx");
    }

    public boolean checkAuth() {

        try {

            ctx = new InitialDirContext(env);

            bLogin = true;

        } catch (javax.naming.AuthenticationException authe) {

            bLogin = false;

        } catch (Exception e) {
        } finally {
            try {
                ctx.close();

            } catch (Exception Ignore) {
            }
        }
        return bLogin;
    }

    public static void main(String[] args) {
        if(args.length==1 || args[1] == ""){
            System.out.print("false");
            return;
        }
        LDAPUtil lDAPUtil = new LDAPUtil(args[0],args[1]);
        boolean auth = lDAPUtil.checkAuth();
        System.out.print(auth);
    }

}
