import {IIconProps} from '../DataStoreIcon';

const Lightstep = ({color, width = '20', height = '20'}: IIconProps) => {
  return (
    <svg width={width} height={height} viewBox="0 0 21 16" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M10.4937 10.5937C11.9183 10.5937 13.0732 9.42379 13.0732 7.98069C13.0732 6.53754 11.9183 5.36768 10.4937 5.36768C9.06917 5.36768 7.91431 6.53754 7.91431 7.98069C7.91431 9.42379 9.06917 10.5937 10.4937 10.5937Z"
        fill={color}
      />
      <path
        d="M9.78891 0.0201156L1.37999 4.58417C1.3284 4.60161 1.31121 4.65387 1.29401 4.70614L0.00429871 10.5593C-0.0128967 10.6464 0.021496 10.7161 0.107477 10.7509L11.0098 15.9595C11.0786 15.9944 11.1646 15.9944 11.2334 15.9595L18.2323 12.7368C18.301 12.702 18.3526 12.6323 18.3697 12.5626L20.932 0.873702C20.9493 0.786603 20.8976 0.68208 20.7944 0.66466H20.7772L9.92646 0.0201156C9.87488 -0.0147252 9.82329 0.00269435 9.78891 0.0201156ZM11.0442 15.6285L0.84691 10.7683H4.63006L13.0218 11.9877L11.0442 15.6285ZM13.3485 12.0226L17.6476 12.6497L11.4569 15.4892L13.3485 12.0226ZM4.87081 10.4722L8.06928 4.09642L17.9228 2.94669L13.1766 11.6742L4.87081 10.4722ZM18.1462 2.61572L8.24125 3.76544L9.97808 0.316257L20.3473 0.943378L18.1462 2.61572ZM7.86294 3.8177L2.2226 4.47967L9.51378 0.525298L7.86294 3.8177ZM1.56915 4.86289L7.70819 4.14869L4.5441 10.4373H0.331027L1.56915 4.86289ZM18.0774 12.3884L13.5033 11.7264L18.3183 2.8596L20.5365 1.16984L18.0774 12.3884Z"
        fill={color}
      />
    </svg>
  );
};

export default Lightstep;