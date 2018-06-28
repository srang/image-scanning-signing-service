FROM registry.access.redhat.com/rhel7/rhel

ENV OCP_VERSION=3.9

ADD bin/sign-image /usr/local/bin/

RUN yum-config-manager --disable rhel-7-server-aus-rpms,rhel-7-server-htb-rpms,rhel-7-server-nfv-rpms,rhel-7-server-rt-rpms,rhel-7-server-tus-rpms && \
    yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm && \
    yum-config-manager --enable rhel-7-server-extras-rpms,rhel-7-server-ose-${OCP_VERSION}-rpms,rhel-7-server-optional-rpms && \
    yum repolist  && \
    yum clean all && \
    INSTALL_PKGS="docker atomic jq atomic-openshift-clients image-inspector tar" && \
    yum install -y --setopt=tsflags=nodocs $INSTALL_PKGS && \
    rpm -V $INSTALL_PKGS && \
    yum clean all