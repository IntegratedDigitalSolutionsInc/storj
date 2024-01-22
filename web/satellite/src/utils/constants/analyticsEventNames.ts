// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

// Make sure these event names match up with the client-side event names in satellite/analytics/service.go
export enum AnalyticsEvent {
    GATEWAY_CREDENTIALS_CREATED = 'Credentials Created',
    PATH_SELECTED = 'Path Selected',
    LINK_SHARED = 'Link Shared',
    OBJECT_UPLOADED = 'Object Uploaded',
    API_KEY_GENERATED = 'API Key Generated',
    UPGRADE_BANNER_CLICKED = 'Upgrade Banner Clicked',
    MODAL_ADD_CARD = 'Credit Card Added In Modal',
    SEARCH_BUCKETS = 'Search Buckets',
    NAVIGATE_PROJECTS = 'Navigate Projects',
    MANAGE_PROJECTS_CLICKED = 'Manage Projects Clicked',
    CREATE_NEW_CLICKED = 'Create New Clicked',
    VIEW_DOCS_CLICKED = 'View Docs Clicked',
    VIEW_FORUM_CLICKED = 'View Forum Clicked',
    VIEW_SUPPORT_CLICKED = 'View Support Clicked',
    CREATE_AN_ACCESS_GRANT_CLICKED = 'Create an Access Grant Clicked',
    UPLOAD_USING_CLI_CLICKED = 'Upload Using CLI Clicked',
    UPLOAD_IN_WEB_CLICKED = 'Upload In Web Clicked',
    NEW_PROJECT_CLICKED = 'New Project Clicked',
    LOGOUT_CLICKED = 'Logout Clicked',
    PROFILE_UPDATED = 'Profile Updated',
    PASSWORD_CHANGED = 'Password Changed',
    MFA_ENABLED = 'MFA Enabled',
    BUCKET_CREATED = 'Bucket Created',
    BUCKET_DELETED = 'Bucket Deleted',
    ACCESS_GRANT_CREATED = 'Access Grant Created',
    API_ACCESS_CREATED  = 'API Access Created',
    UPLOAD_FILE_CLICKED = 'Upload File Clicked',
    UPLOAD_FOLDER_CLICKED = 'Upload Folder Clicked',
    DOWNLOAD_TXT_CLICKED = 'Download txt clicked',
    COPY_TO_CLIPBOARD_CLICKED = 'Copy to Clipboard Clicked',
    CREATE_ACCESS_GRANT_CLICKED = 'Create Access Grant Clicked',
    CREATE_S3_CREDENTIALS_CLICKED = 'Create S3 Credentials Clicked',
    CREATE_KEYS_FOR_CLI_CLICKED = 'Create Keys For CLI Clicked',
    SEE_PAYMENTS_CLICKED = 'See Payments Clicked',
    EDIT_PAYMENT_METHOD_CLICKED = 'Edit Payment Method Clicked',
    ADD_NEW_PAYMENT_METHOD_CLICKED = 'Add New Payment Method Clicked',
    APPLY_NEW_COUPON_CLICKED = 'Apply New Coupon Clicked',
    CREDIT_CARD_REMOVED = 'Credit Card Removed',
    COUPON_CODE_APPLIED = 'Coupon Code Applied',
    INVOICE_DOWNLOADED = 'Invoice Downloaded',
    CREDIT_CARD_ADDED_FROM_BILLING = 'Credit Card Added From Billing',
    ADD_FUNDS_CLICKED = 'Add Funds Clicked',
    PROJECT_MEMBERS_INVITE_SENT = 'Project Members Invite Sent',
    UI_ERROR = 'UI error occurred',
    PROJECT_NAME_UPDATED = 'Project Name Updated',
    PROJECT_DESCRIPTION_UPDATED = 'Project Description Updated',
    PROJECT_STORAGE_LIMIT_UPDATED = 'Project Storage Limit Updated',
    PROJECT_BANDWIDTH_LIMIT_UPDATED = 'Project Bandwidth Limit Updated',
    GALLERY_VIEW_CLICKED = 'Gallery View Clicked',
    PROJECT_INVITATION_ACCEPTED = 'Project Invitation Accepted',
    PROJECT_INVITATION_DECLINED = 'Project Invitation Declined',
    PASSPHRASE_CREATED = 'Passphrase Created',
    RESEND_INVITE_CLICKED = 'Resend Invite Clicked',
    COPY_INVITE_LINK_CLICKED = 'Copy Invite Link Clicked',
    REMOVE_PROJECT_MEMBER_CLICKED = 'Remove Member Clicked',
}

export enum AnalyticsErrorEventSource {
    ACCESS_GRANTS_PAGE = 'Access grants page',
    ACCOUNT_PAGE = 'Account page',
    ACCOUNT_SETTINGS_AREA = 'Account settings area',
    ACCOUNT_SETUP_DIALOG = 'Account setup dialog',
    BILLING_HISTORY_TAB = 'Billing history tab',
    BILLING_COUPONS_TAB = 'Billing coupons tab',
    BILLING_OVERVIEW_TAB = 'Billing overview tab',
    BILLING_PAYMENT_METHODS_TAB = 'Billing payment methods tab',
    BILLING_APPLY_COUPON_CODE_INPUT = 'Billing apply coupon code input',
    BILLING_STRIPE_CARD_INPUT = 'Billing stripe card input',
    BILLING_AREA = 'Billing area',
    BILLING_STORJ_TOKEN_CONTAINER = 'Billing STORJ token container',
    CREATE_AG_MODAL = 'Create access grant modal',
    CONFIRM_DELETE_AG_MODAL = 'Confirm delete access grant modal',
    FILE_BROWSER_LIST_CALL = 'File browser - list API call',
    FILE_BROWSER_ENTRY = 'File browser entry',
    FILE_BROWSER = 'File browser',
    PROJECT_INFO_BAR = 'Project info bar',
    UPGRADE_ACCOUNT_MODAL = 'Upgrade account modal',
    ADD_PROJECT_MEMBER_MODAL = 'Add project member modal',
    ADD_TOKEN_FUNDS_MODAL = 'Add token funds modal',
    CHANGE_PROJECT_LIMIT_MODAL = 'Change project limit modal',
    REQUEST_PROJECT_LIMIT_MODAL = 'Request project limit modal',
    CHANGE_PASSWORD_MODAL = 'Change password modal',
    CREATE_PROJECT_MODAL = 'Create project modal',
    CREATE_PROJECT_PASSPHRASE_MODAL = 'Create project passphrase modal',
    CREATE_BUCKET_MODAL = 'Create bucket modal',
    DELETE_BUCKET_MODAL = 'Delete bucket modal',
    ENABLE_MFA_MODAL = 'Enable MFA modal',
    MFA_CODES_MODAL = 'MFA codes modal',
    DISABLE_MFA_MODAL = 'Disable MFA modal',
    EDIT_PROFILE_MODAL = 'Edit profile modal',
    CREATE_FOLDER_MODAL = 'Create folder modal',
    OBJECT_DETAILS_MODAL = 'Object details modal',
    OPEN_BUCKET_MODAL = 'Open bucket modal',
    SHARE_MODAL = 'Share modal',
    OBJECTS_UPLOAD_MODAL = 'Objects upload modal',
    NAVIGATION_ACCOUNT_AREA = 'Navigation account area',
    NAVIGATION_PROJECT_SELECTION = 'Navigation project selection',
    MOBILE_NAVIGATION = 'Mobile navigation',
    BUCKET_TABLE = 'Bucket table',
    BUCKET_PAGE = 'Bucket page',
    BUCKET_DETAILS_PAGE = 'Bucket details page',
    UPLOAD_FILE_VIEW = 'Upload file view',
    GALLERY_VIEW = 'Gallery view',
    OBJECT_UPLOAD_ERROR = 'Object upload error',
    ONBOARDING_NAME_STEP = 'Onboarding name step',
    ONBOARDING_PERMISSIONS_STEP = 'Onboarding permissions step',
    PROJECT_DASHBOARD_PAGE = 'Project dashboard page',
    PROJECT_SETTINGS_AREA = 'Project settings area',
    EDIT_PROJECT_DETAILS = 'Edit project details',
    EDIT_PROJECT_LIMIT = 'Edit project limit',
    PROJECTS_LIST = 'Projects list',
    PROJECT_MEMBERS_HEADER = 'Project members page header',
    PROJECT_MEMBERS_PAGE = 'Project members page',
    OVERALL_APP_WRAPPER_ERROR = 'Overall app wrapper error',
    OVERALL_SESSION_EXPIRED_ERROR = 'Overall session expired error',
    ALL_PROJECT_DASHBOARD = 'All projects dashboard error',
    ONBOARDING_OVERVIEW_STEP = 'Onboarding Overview step error',
    PRICING_PLAN_STEP = 'Onboarding Pricing Plan step error',
    EDIT_TIMEOUT_MODAL = 'Edit session timeout error',
    SKIP_PASSPHRASE_MODAL = 'Remember skip passphrase error',
    JOIN_PROJECT_MODAL = 'Join project modal',
    PROJECT_INVITATION = 'Project invitation',
    DETAILED_USAGE_REPORT_MODAL = 'Detailed usage report modal',
    REMOVE_CC_MODAL = 'Remove credit card modal',
    EDIT_DEFAULT_CC_MODAL = 'Edit default credit card modal',
}
