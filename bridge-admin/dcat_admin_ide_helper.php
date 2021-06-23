<?php

/**
 * A helper file for Dcat Admin, to provide autocomplete information to your IDE
 *
 * This file should not be included in your code, only analyzed by your IDE!
 *
 * @author jqh <841324345@qq.com>
 */
namespace Dcat\Admin {
    use Illuminate\Support\Collection;

    /**
     * @property Grid\Column|Collection created_at
     * @property Grid\Column|Collection detail
     * @property Grid\Column|Collection id
     * @property Grid\Column|Collection name
     * @property Grid\Column|Collection type
     * @property Grid\Column|Collection updated_at
     * @property Grid\Column|Collection version
     * @property Grid\Column|Collection is_enabled
     * @property Grid\Column|Collection extension
     * @property Grid\Column|Collection icon
     * @property Grid\Column|Collection order
     * @property Grid\Column|Collection parent_id
     * @property Grid\Column|Collection uri
     * @property Grid\Column|Collection menu_id
     * @property Grid\Column|Collection permission_id
     * @property Grid\Column|Collection http_method
     * @property Grid\Column|Collection http_path
     * @property Grid\Column|Collection slug
     * @property Grid\Column|Collection role_id
     * @property Grid\Column|Collection user_id
     * @property Grid\Column|Collection value
     * @property Grid\Column|Collection avatar
     * @property Grid\Column|Collection password
     * @property Grid\Column|Collection remember_token
     * @property Grid\Column|Collection username
     * @property Grid\Column|Collection balance
     * @property Grid\Column|Collection chainId
     * @property Grid\Column|Collection token
     * @property Grid\Column|Collection tokenName
     * @property Grid\Column|Collection bridge
     * @property Grid\Column|Collection sort
     * @property Grid\Column|Collection status
     * @property Grid\Column|Collection url
     * @property Grid\Column|Collection connection
     * @property Grid\Column|Collection exception
     * @property Grid\Column|Collection failed_at
     * @property Grid\Column|Collection payload
     * @property Grid\Column|Collection queue
     * @property Grid\Column|Collection uuid
     * @property Grid\Column|Collection depositHash
     * @property Grid\Column|Collection depositTime
     * @property Grid\Column|Collection fromChain
     * @property Grid\Column|Collection pairId
     * @property Grid\Column|Collection recipient
     * @property Grid\Column|Collection toChain
     * @property Grid\Column|Collection withdrawHash
     * @property Grid\Column|Collection withdrawTime
     * @property Grid\Column|Collection bridgeFee
     * @property Grid\Column|Collection decimal
     * @property Grid\Column|Collection fromToken
     * @property Grid\Column|Collection isMain
     * @property Grid\Column|Collection isNative
     * @property Grid\Column|Collection limit
     * @property Grid\Column|Collection minValue
     * @property Grid\Column|Collection tokenFee
     * @property Grid\Column|Collection toToken
     * @property Grid\Column|Collection email
     * @property Grid\Column|Collection email_verified_at
     *
     * @method Grid\Column|Collection created_at(string $label = null)
     * @method Grid\Column|Collection detail(string $label = null)
     * @method Grid\Column|Collection id(string $label = null)
     * @method Grid\Column|Collection name(string $label = null)
     * @method Grid\Column|Collection type(string $label = null)
     * @method Grid\Column|Collection updated_at(string $label = null)
     * @method Grid\Column|Collection version(string $label = null)
     * @method Grid\Column|Collection is_enabled(string $label = null)
     * @method Grid\Column|Collection extension(string $label = null)
     * @method Grid\Column|Collection icon(string $label = null)
     * @method Grid\Column|Collection order(string $label = null)
     * @method Grid\Column|Collection parent_id(string $label = null)
     * @method Grid\Column|Collection uri(string $label = null)
     * @method Grid\Column|Collection menu_id(string $label = null)
     * @method Grid\Column|Collection permission_id(string $label = null)
     * @method Grid\Column|Collection http_method(string $label = null)
     * @method Grid\Column|Collection http_path(string $label = null)
     * @method Grid\Column|Collection slug(string $label = null)
     * @method Grid\Column|Collection role_id(string $label = null)
     * @method Grid\Column|Collection user_id(string $label = null)
     * @method Grid\Column|Collection value(string $label = null)
     * @method Grid\Column|Collection avatar(string $label = null)
     * @method Grid\Column|Collection password(string $label = null)
     * @method Grid\Column|Collection remember_token(string $label = null)
     * @method Grid\Column|Collection username(string $label = null)
     * @method Grid\Column|Collection balance(string $label = null)
     * @method Grid\Column|Collection chainId(string $label = null)
     * @method Grid\Column|Collection token(string $label = null)
     * @method Grid\Column|Collection tokenName(string $label = null)
     * @method Grid\Column|Collection bridge(string $label = null)
     * @method Grid\Column|Collection sort(string $label = null)
     * @method Grid\Column|Collection status(string $label = null)
     * @method Grid\Column|Collection url(string $label = null)
     * @method Grid\Column|Collection connection(string $label = null)
     * @method Grid\Column|Collection exception(string $label = null)
     * @method Grid\Column|Collection failed_at(string $label = null)
     * @method Grid\Column|Collection payload(string $label = null)
     * @method Grid\Column|Collection queue(string $label = null)
     * @method Grid\Column|Collection uuid(string $label = null)
     * @method Grid\Column|Collection depositHash(string $label = null)
     * @method Grid\Column|Collection depositTime(string $label = null)
     * @method Grid\Column|Collection fromChain(string $label = null)
     * @method Grid\Column|Collection pairId(string $label = null)
     * @method Grid\Column|Collection recipient(string $label = null)
     * @method Grid\Column|Collection toChain(string $label = null)
     * @method Grid\Column|Collection withdrawHash(string $label = null)
     * @method Grid\Column|Collection withdrawTime(string $label = null)
     * @method Grid\Column|Collection bridgeFee(string $label = null)
     * @method Grid\Column|Collection decimal(string $label = null)
     * @method Grid\Column|Collection fromToken(string $label = null)
     * @method Grid\Column|Collection isMain(string $label = null)
     * @method Grid\Column|Collection isNative(string $label = null)
     * @method Grid\Column|Collection limit(string $label = null)
     * @method Grid\Column|Collection minValue(string $label = null)
     * @method Grid\Column|Collection tokenFee(string $label = null)
     * @method Grid\Column|Collection toToken(string $label = null)
     * @method Grid\Column|Collection email(string $label = null)
     * @method Grid\Column|Collection email_verified_at(string $label = null)
     */
    class Grid {}

    class MiniGrid extends Grid {}

    /**
     * @property Show\Field|Collection created_at
     * @property Show\Field|Collection detail
     * @property Show\Field|Collection id
     * @property Show\Field|Collection name
     * @property Show\Field|Collection type
     * @property Show\Field|Collection updated_at
     * @property Show\Field|Collection version
     * @property Show\Field|Collection is_enabled
     * @property Show\Field|Collection extension
     * @property Show\Field|Collection icon
     * @property Show\Field|Collection order
     * @property Show\Field|Collection parent_id
     * @property Show\Field|Collection uri
     * @property Show\Field|Collection menu_id
     * @property Show\Field|Collection permission_id
     * @property Show\Field|Collection http_method
     * @property Show\Field|Collection http_path
     * @property Show\Field|Collection slug
     * @property Show\Field|Collection role_id
     * @property Show\Field|Collection user_id
     * @property Show\Field|Collection value
     * @property Show\Field|Collection avatar
     * @property Show\Field|Collection password
     * @property Show\Field|Collection remember_token
     * @property Show\Field|Collection username
     * @property Show\Field|Collection balance
     * @property Show\Field|Collection chainId
     * @property Show\Field|Collection token
     * @property Show\Field|Collection tokenName
     * @property Show\Field|Collection bridge
     * @property Show\Field|Collection sort
     * @property Show\Field|Collection status
     * @property Show\Field|Collection url
     * @property Show\Field|Collection connection
     * @property Show\Field|Collection exception
     * @property Show\Field|Collection failed_at
     * @property Show\Field|Collection payload
     * @property Show\Field|Collection queue
     * @property Show\Field|Collection uuid
     * @property Show\Field|Collection depositHash
     * @property Show\Field|Collection depositTime
     * @property Show\Field|Collection fromChain
     * @property Show\Field|Collection pairId
     * @property Show\Field|Collection recipient
     * @property Show\Field|Collection toChain
     * @property Show\Field|Collection withdrawHash
     * @property Show\Field|Collection withdrawTime
     * @property Show\Field|Collection bridgeFee
     * @property Show\Field|Collection decimal
     * @property Show\Field|Collection fromToken
     * @property Show\Field|Collection isMain
     * @property Show\Field|Collection isNative
     * @property Show\Field|Collection limit
     * @property Show\Field|Collection minValue
     * @property Show\Field|Collection tokenFee
     * @property Show\Field|Collection toToken
     * @property Show\Field|Collection email
     * @property Show\Field|Collection email_verified_at
     *
     * @method Show\Field|Collection created_at(string $label = null)
     * @method Show\Field|Collection detail(string $label = null)
     * @method Show\Field|Collection id(string $label = null)
     * @method Show\Field|Collection name(string $label = null)
     * @method Show\Field|Collection type(string $label = null)
     * @method Show\Field|Collection updated_at(string $label = null)
     * @method Show\Field|Collection version(string $label = null)
     * @method Show\Field|Collection is_enabled(string $label = null)
     * @method Show\Field|Collection extension(string $label = null)
     * @method Show\Field|Collection icon(string $label = null)
     * @method Show\Field|Collection order(string $label = null)
     * @method Show\Field|Collection parent_id(string $label = null)
     * @method Show\Field|Collection uri(string $label = null)
     * @method Show\Field|Collection menu_id(string $label = null)
     * @method Show\Field|Collection permission_id(string $label = null)
     * @method Show\Field|Collection http_method(string $label = null)
     * @method Show\Field|Collection http_path(string $label = null)
     * @method Show\Field|Collection slug(string $label = null)
     * @method Show\Field|Collection role_id(string $label = null)
     * @method Show\Field|Collection user_id(string $label = null)
     * @method Show\Field|Collection value(string $label = null)
     * @method Show\Field|Collection avatar(string $label = null)
     * @method Show\Field|Collection password(string $label = null)
     * @method Show\Field|Collection remember_token(string $label = null)
     * @method Show\Field|Collection username(string $label = null)
     * @method Show\Field|Collection balance(string $label = null)
     * @method Show\Field|Collection chainId(string $label = null)
     * @method Show\Field|Collection token(string $label = null)
     * @method Show\Field|Collection tokenName(string $label = null)
     * @method Show\Field|Collection bridge(string $label = null)
     * @method Show\Field|Collection sort(string $label = null)
     * @method Show\Field|Collection status(string $label = null)
     * @method Show\Field|Collection url(string $label = null)
     * @method Show\Field|Collection connection(string $label = null)
     * @method Show\Field|Collection exception(string $label = null)
     * @method Show\Field|Collection failed_at(string $label = null)
     * @method Show\Field|Collection payload(string $label = null)
     * @method Show\Field|Collection queue(string $label = null)
     * @method Show\Field|Collection uuid(string $label = null)
     * @method Show\Field|Collection depositHash(string $label = null)
     * @method Show\Field|Collection depositTime(string $label = null)
     * @method Show\Field|Collection fromChain(string $label = null)
     * @method Show\Field|Collection pairId(string $label = null)
     * @method Show\Field|Collection recipient(string $label = null)
     * @method Show\Field|Collection toChain(string $label = null)
     * @method Show\Field|Collection withdrawHash(string $label = null)
     * @method Show\Field|Collection withdrawTime(string $label = null)
     * @method Show\Field|Collection bridgeFee(string $label = null)
     * @method Show\Field|Collection decimal(string $label = null)
     * @method Show\Field|Collection fromToken(string $label = null)
     * @method Show\Field|Collection isMain(string $label = null)
     * @method Show\Field|Collection isNative(string $label = null)
     * @method Show\Field|Collection limit(string $label = null)
     * @method Show\Field|Collection minValue(string $label = null)
     * @method Show\Field|Collection tokenFee(string $label = null)
     * @method Show\Field|Collection toToken(string $label = null)
     * @method Show\Field|Collection email(string $label = null)
     * @method Show\Field|Collection email_verified_at(string $label = null)
     */
    class Show {}

    /**
     
     */
    class Form {}

}

namespace Dcat\Admin\Grid {
    /**
     
     */
    class Column {}

    /**
     
     */
    class Filter {}
}

namespace Dcat\Admin\Show {
    /**
     
     */
    class Field {}
}
