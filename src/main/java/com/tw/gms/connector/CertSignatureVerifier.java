package com.tw.gms.connector;

import org.bouncycastle.asn1.ASN1InputStream;
import org.bouncycastle.asn1.ASN1ObjectIdentifier;
import org.bouncycastle.asn1.DEROctetString;
import org.bouncycastle.asn1.DLSequence;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;
import java.io.IOException;
import java.security.*;
import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.Arrays;
import java.util.Enumeration;

@Service
public class CertSignatureVerifier {
    Logger log = LoggerFactory.getLogger(CertSignatureVerifier.class);

    public boolean verifyCertChainSignatures(X509Certificate[] certChain) throws Exception {
        return verifyCertChainSignatures2(certChain);
        //return verifyCertChainSignatures1(certChain);
    }

    private boolean verifyCertChainSignatures2(X509Certificate[] x509Certificates) throws CertificateException {
        int n = x509Certificates.length;
        for (int i = 0; i < n - 1; i++) {
            X509Certificate cert = x509Certificates[i];
            X509Certificate issuer = x509Certificates[i + 1];
            if (!cert.getIssuerX500Principal().equals(issuer.getSubjectX500Principal())) {
                throw new RuntimeException("Certificates do not chain");
            }
            try {
                cert.verify(issuer.getPublicKey());
            } catch (InvalidKeyException | NoSuchProviderException | SignatureException |
                     NoSuchAlgorithmException e) {
                throw new RuntimeException(e);
            }
            log.info("Verified: {}", cert.getSubjectX500Principal());
        }
        X509Certificate last = x509Certificates[n - 1];
        // if self-signed, verify the final cert
        if (last.getIssuerX500Principal().equals(last.getSubjectX500Principal())) {
            try {
                last.verify(last.getPublicKey());
            } catch (InvalidKeyException | NoSuchProviderException | SignatureException |
                     NoSuchAlgorithmException e) {
                throw new RuntimeException(e);
            }
            log.info("Verified: {}", last.getSubjectX500Principal());
        }
        return true;
    }

    private boolean verifyCertChainSignatures1(X509Certificate[] certChain) throws IllegalBlockSizeException, BadPaddingException, NoSuchPaddingException, InvalidKeyException, CertificateEncodingException, SignatureException, IOException, NoSuchAlgorithmException {
        boolean signatureChainIsValid = true;
        for (int i = certChain.length - 1; i > 0; i--) {
            signatureChainIsValid = verifyCertSignature(certChain[i], certChain[i - 1]);
        }
        // Check whether root certificate is signed by itself
        X509Certificate rootCert = certChain[0];
        signatureChainIsValid = verifyCertSignature(rootCert, rootCert);
        return signatureChainIsValid;
    }

    public boolean verifyCertSignature(X509Certificate lowerCert, X509Certificate higherCert)
            throws IllegalBlockSizeException, BadPaddingException, NoSuchPaddingException, InvalidKeyException,
            CertificateEncodingException, SignatureException, IOException, NoSuchAlgorithmException {
        // Compute certificate digest
        String hashAlgName = lowerCert.getSigAlgName().split("with")[0];
        MessageDigest md = MessageDigest.getInstance(hashAlgName);
        byte[] tbsCertificate = lowerCert.getTBSCertificate();
        md.update(tbsCertificate);
        byte[] computedDigest = md.digest();
        // Decode signature
        String cryptoAlgName = lowerCert.getSigAlgName().split("with")[1];
        Cipher decCipher = Cipher.getInstance(cryptoAlgName);
        decCipher.init(Cipher.DECRYPT_MODE, higherCert.getPublicKey());
        byte[] decodedSignature = decCipher.doFinal(lowerCert.getSignature());
        byte[] decodedDigest = extractAsn1EncodedSignature(decodedSignature);
        return Arrays.equals(computedDigest, decodedDigest);
    }

    private byte[] extractAsn1EncodedSignature(byte[] bytes) throws IOException {
        ASN1InputStream ais = new ASN1InputStream(bytes);
        DLSequence superSeq = (DLSequence) ais.readObject();
        // Extract signature bytes
        Enumeration e1 = superSeq.getObjects();
        DLSequence subSeq = (DLSequence) e1.nextElement();
        DEROctetString derOctetString = (DEROctetString) e1.nextElement();
        byte[] octets = derOctetString.getOctets();
        // Extract signature algorithm OID string (not used here, though)
        Enumeration e2 = subSeq.getObjects();
        ASN1ObjectIdentifier algorithmIdentifier = (ASN1ObjectIdentifier) e2.nextElement();
        String oidString = algorithmIdentifier.toString();
        return octets;
    }
}
